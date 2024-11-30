import { Service as Spotify } from "$lib/server/spotify"
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
    const { spotify, accessToken } = await Spotify(p.user.accessToken, p.user.refreshToken)

    await db
      .update(usersTable)
      .set({ accessToken })
      .where(eq(usersTable.id, p.user.id))


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
