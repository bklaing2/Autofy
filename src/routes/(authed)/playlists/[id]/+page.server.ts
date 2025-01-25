import type { Actions } from './$types';
import type { PageServerLoad } from './$types'
import { error, redirect } from "@sveltejs/kit";
import Playlist from "$lib/server/playlist.js"
import { eq } from 'drizzle-orm'
import { playlistsTable } from "$lib/server/db/schema";

export const load: PageServerLoad = async ({ params, locals }) => {
  const { db, spotify } = locals
  const playlistId = params.id

  const [dbPlaylist] = await db
    .select()
    .from(playlistsTable)
    .where(eq(playlistsTable.id, playlistId))
    .limit(1)

  if (!dbPlaylist) throw error(404, { message: 'Playlist not found' })

  const { body: spotifyPlaylist } = await spotify.getPlaylist(playlistId)

  const json = {
    ...dbPlaylist,
    id: playlistId,
    title: spotifyPlaylist.name,
    artists: await Playlist.getArtists(dbPlaylist.artists, spotify),
    followedArtists: dbPlaylist.followedArtists.length > 0
  }
  console.log(dbPlaylist.followedArtists)
  console.log(json.followedArtists)
  return { playlist: json }
}


export const actions = {
  update: async ({ params, request, fetch }) => {
    await fetch(`/playlists/${params.id}`, { method: 'PATCH', body: await request.formData() })
    throw redirect(303, '/playlists')
  }
} satisfies Actions
