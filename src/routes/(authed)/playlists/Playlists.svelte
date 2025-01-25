<script lang="ts">
	import { page } from '$app/stores';
	import Playlist from './Playlist.svelte';

	export let playlists: string[];
</script>

<svelte:head>
	<title>my playlists</title>
	<meta name="description" content="My autofy playlists" />
</svelte:head>

<main class="layout">
	<h1>my autofy playlists</h1>

	{#each playlists as playlist}
		<Playlist id={playlist} flipped={$page.params.id === playlist}>
			{#if $page.params.id === playlist}
				<slot />
			{/if}
		</Playlist>
	{/each}

	<a class="playlist wireframe" href="/create-playlist">+</a>
</main>

<style>
	.layout {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		column-gap: 2rem;
		row-gap: 3rem;
	}

	h1 {
		width: 100%;
		font-size: 2rem;
		text-align: center;
	}

	.playlist {
		max-width: 100%;
		max-height: 100%;
		aspect-ratio: 1 / 1.414;
	}

	.wireframe {
		width: 300px;

		font-size: 5rem;
		color: gray;
		border: 2px dashed gray;
		border-radius: 12px;

		display: flex;
		justify-content: center;
		align-items: center;
	}
</style>

