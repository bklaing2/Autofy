import { error, redirect } from "@sveltejs/kit";
import Playlist from "$lib/server/playlist.js"

export async function load({ params, locals }) {
  const { supabase, spotify } = locals
  const playlistId = params.id

  const { data: supabasePlaylist, error: supabaseError } = await supabase
    .from('playlists')
    .select('*')
    .eq('id', playlistId)
    .limit(1)
    .single()

  if (supabaseError) throw new Error (supabaseError.message)
  if (!supabasePlaylist) throw error(404, { message: 'Playlist not found' })

  const { body: spotifyPlaylist } = await spotify.getPlaylist(playlistId)

  const json = {
    ...supabasePlaylist,
    id: playlistId,
    title: spotifyPlaylist.name,
    artists: await Playlist.getArtists(supabasePlaylist.artists, spotify),
    followed_artists: !!supabasePlaylist.followed_artists
  }
  return { playlist: json }
}


export const actions = {
  update: async ({ params, request, fetch }) => {
    await fetch(`/playlists/${params.id}`, { method: 'PATCH', body: await request.formData() })
    throw redirect(303, '/playlists')
  }
}