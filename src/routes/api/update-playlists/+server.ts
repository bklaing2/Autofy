import { Service as Spotify } from "$lib/server/spotify"
import Playlist from "$lib/server/playlist"
import { eq } from 'drizzle-orm'
import db from "$lib/server/db"
import { playlistsTable, tokensTable } from "$lib/server/db/schema"

export async function GET() {

  const data = await db
    .select()
    .from(playlistsTable)
    .innerJoin(tokensTable, eq(playlistsTable.userId, tokensTable.userId))

  await Promise.all(data.map(async d => {
    if (!d.tokens) return
    const { spotify, accessToken } = await Spotify(d.tokens?.accessToken, d.tokens?.refreshToken)

    await db
      .update(tokensTable)
      .set({ accessToken })
      .where(eq(tokensTable.userId, d.tokens.userId))



    // Delete if the user removed the playlist
    const exists = await spotify.getUserPlaylists()
    if (!exists.body.items.find(p => p.id === d.playlists.id)) {
      await db
        .delete(playlistsTable)
        .where(eq(playlistsTable.id, d.playlists.id))
      return
    }

    Playlist.update(d.playlists, spotify)
  }))

  return new Response(null, { status: 204 })
}
