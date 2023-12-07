import Playlist from "$lib/server/playlist.js"


export async function PATCH({ params, request, locals }) {
	const { supabase, spotify } = locals
	const playlistId = params.id

	const form = await request.formData()
	const artists = [ ...new Set(form.getAll('artist') as string[]) ]
	const includeFollowedArtists = form.has('followed-artists')

	
	const title = form.get('title') as string || 'autofy playlist'
	const now = new Date()
	const date = `${now.getMonth() + 1}/${now.getDate()}/${now.getFullYear().toString().slice(-2)}`
	
	await spotify.changePlaylistDetails(playlistId, {
		name: title,
		description: `Updated on ${date} by autofy`
	})

	const updateWhen = form.getAll('update-when') as string[]

	const followed = includeFollowedArtists ? await Playlist.getFollowedArtists(spotify) : null
	console.log(`Artists: ${artists.length}`)


	const { data: oldPlaylist } = await supabase
		.from('playlists')
		.select('*')
		.eq('id', playlistId)
		.limit(1)
		.single()

	const { data: newPlaylist } = await supabase
		.from('playlists')
		.upsert({
			id: playlistId,
			artists: artists,
			followed_artists: followed,
			update_when_artist_posts: updateWhen.includes('artist-posts'),
			update_when_user_follows_artist: updateWhen.includes('user-follows-artist'),
			update_when_user_unfollows_artist: updateWhen.includes('user-unfollows-artist')
		})
		.eq('id', playlistId)
		.select('*')
		.limit(1)
		.single()

	
	Playlist.updateSettings(oldPlaylist, newPlaylist, spotify)
	return new Response(null, { status: 204 })
}
