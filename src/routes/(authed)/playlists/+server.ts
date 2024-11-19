import type { RequestHandler } from './$types';
import { redirect } from '@sveltejs/kit'
import Playlist from '$lib/server/playlist'
import { playlistImageBase64 } from '$lib/server/const.js'
import { playlistsTable } from '$lib/server/db/schema';


export const POST: RequestHandler = async ({ request, locals }) => {
  const { db, spotify } = locals

  const form = await request.formData()
  const artists = [...new Set(form.getAll('artist') as string[])]
  const includeFollowedArtists = form.has('followed-artists')

  if (artists.length === 0 && !includeFollowedArtists) return new Response(null, { status: 204 })


  const title = form.get('title') as string || 'new autofy playlist'
  const now = new Date()
  const date = `${now.getMonth() + 1}/${now.getDate()}/${now.getFullYear().toString().slice(-2)}`

  const data = await spotify.createPlaylist(title, {
    description: `Updated on ${date} by autofy`,
    public: true
  })

  const playlistId = data.body.id

  try {
    await spotify.uploadCustomPlaylistCoverImage(playlistId, playlistImageBase64)
  } catch (e) {
    console.log(e)
  }


  const updateWhen = form.getAll('update-when') as string[]

  const followedArtists = includeFollowedArtists ? await Playlist.getFollowedArtists(spotify) : undefined
  console.log(`Artists: ${artists.length}`)


  await db
    .insert(playlistsTable)
    .values({
      id: playlistId,
      userId: (await spotify.getMe()).body.id,
      artists,
      followedArtists,
      updateWhenArtistPosts: updateWhen.includes('artist-posts'),
      updateWhenUserFollowsArtist: updateWhen.includes('user-follows-artist'),
      updateWhenUserUnfollowsArtist: updateWhen.includes('user-unfollows-artist')
    })



  Playlist.populate(data.body.id, (followedArtists ?? []).concat(artists), spotify)
  throw redirect(303, '/playlists')
}
