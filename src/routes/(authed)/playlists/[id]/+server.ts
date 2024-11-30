import type { RequestHandler } from './$types';
import Playlist from "$lib/server/playlist.js"
import { eq } from 'drizzle-orm'
import { playlistsTable } from '$lib/server/db/schema';


export const PATCH: RequestHandler = async ({ params, request, locals }) => {
  const { db, spotify } = locals
  const playlistId = params.id

  const form = await request.formData()
  const artists = [...new Set(form.getAll('artist') as string[])]
  const includeFollowedArtists = form.has('followed-artists')


  const title = form.get('title') as string || 'autofy playlist'
  const now = new Date()
  const date = `${now.getMonth() + 1}/${now.getDate()}/${now.getFullYear().toString().slice(-2)}`

  await spotify.changePlaylistDetails(playlistId, {
    name: title,
    description: `Updated on ${date} by autofy`
  })

  const updateWhen = form.getAll('update-when') as string[]

  const followedArtists = includeFollowedArtists ? await Playlist.getFollowedArtists(spotify) : []
  console.log(`Artists: ${artists.length}`)

  const data = {
    id: playlistId,
    userId: (await spotify.getMe()).body.id,
    artists,
    followedArtists,
    updateWhenArtistPosts: updateWhen.includes('artist-posts'),
    updateWhenUserFollowsArtist: updateWhen.includes('user-follows-artist'),
    updateWhenUserUnfollowsArtist: updateWhen.includes('user-unfollows-artist')
  }


  const [oldPlaylist] = await db
    .select()
    .from(playlistsTable)
    .where(eq(playlistsTable.id, playlistId))
    .limit(1)

  const [newPlaylist] = await db
    .insert(playlistsTable)
    .values(data)
    .onConflictDoUpdate({ target: playlistsTable.id, set: data })
    .returning()

  Playlist.updateSettings(oldPlaylist, newPlaylist, spotify)
  return new Response(null, { status: 204 })
}
