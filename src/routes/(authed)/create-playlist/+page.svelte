<script lang="ts">
	import Playlist from '$lib/CreatePlaylist/Playlist.svelte';

	let type: 'followed artists' | 'select artists' = 'select artists'
</script>


<svelte:head>
	<title>create playlist</title>
	<meta name="description" content="My autofy playlists" />
</svelte:head>


<main class="layout">
	<h1 class="header">create autofy playlist</h1>
	<Playlist action='/create-playlist' />
</main>


<style>
	.layout {
		display: flex;
		flex-direction: column;
		align-items: center;
		row-gap: 4rem;
		padding-left: 2rem;
		padding-right: 2rem;
	}

	.header {
		margin: 0;
		margin-bottom: 2rem;
		font-size: 2.5rem;
		text-align: center;
	}
</style>