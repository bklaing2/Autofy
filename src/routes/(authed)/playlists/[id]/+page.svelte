<script>
	import Playlist from "$lib/EditPlaylist/Playlist.svelte";

	export let data
	const playlist = data.playlist
</script>

<svelte:head>
	<title>Playlist Details</title>
	<meta name="description" content="My autofy playlists" />
</svelte:head>



<Playlist {playlist} />