import { redirect } from '@sveltejs/kit'
import Playlist from '$lib/server/playlist'
import { playlistImageBase64 } from '$lib/server/const.js'


export async function POST({ request, locals }) {
	const { supabase, spotify } = locals

	const form = await request.formData()
	const artists = [ ...new Set(form.getAll('artist') as string[]) ]
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

	const followed = includeFollowedArtists ? await Playlist.getFollowedArtists(spotify) : null
	console.log(`Artists: ${artists.length}`)


	await supabase
		.from('playlists')
		.insert({
			id: playlistId,
			artists: artists,
			followed_artists: followed,
			update_when_artist_posts: updateWhen.includes('artist-posts'),
			update_when_user_follows_artist: updateWhen.includes('user-follows-artist'),
			update_when_user_unfollows_artist: updateWhen.includes('user-unfollows-artist')
		})

	
	
	Playlist.populate(data.body.id, (followed ?? []).concat(artists), spotify)
	throw redirect(303, '/playlists')
}
