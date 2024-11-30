import Spotify from "$lib/server/spotify"
import Playlist from "$lib/server/playlist"
import { eq } from 'drizzle-orm'
import db from "$lib/server/db"
import { playlistsTable, usersTable } from "$lib/server/db/schema"

export async function GET() {

  const playlists = await db.query.playlistsTable.findMany({
    with: {
      user: true
    }
  })


  await Promise.all(playlists.map(async p => {
    const spotify = await Spotify(p.user.accessToken, p.user.refreshToken)
    const accessToken = spotify.getAccessToken()
    const refreshToken = spotify.getRefreshToken()

    if (!accessToken || !refreshToken) return


    await db
      .insert(usersTable)
      .values({ id: p.userId, accessToken, refreshToken })
      .onConflictDoUpdate({ target: usersTable.id, set: { accessToken, refreshToken } });


    // Delete if the user removed the playlist
    const exists = await spotify.getUserPlaylists()
    if (!exists.body.items.find(p => p.id === p.id)) {
      await db
        .delete(playlistsTable)
        .where(eq(playlistsTable.id, p.id))
      return
    }

    Playlist.update(p, spotify)
  }))

  return new Response(null, { status: 204 })
}
