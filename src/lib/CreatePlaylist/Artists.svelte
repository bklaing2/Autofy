<script lang="ts" context="module">
	export type Artist = { id: string; href: string; name: string; img?: string };
</script>

<script lang="ts">
	import Section from './Section.svelte';
	import ArtistModal from '$lib/SearchArtists.svelte';
	import Option from './Option.svelte';

	export let artists: Artist[];
	export let hidden = false;
	export let color = 'gray';

	export let followedArtists = false;

	let addArtist: ArtistModal;

	function artistSelected(a: SpotifyApi.ArtistObjectFull) {
		const artist: Artist = {
			id: a.id,
			href: a.href,
			name: a.name,
			img: a.images && a.images.length > 0 ? a.images[0].url : undefined
		};

		artists = !artists.find((a) => a.id === artist.id) ? [...artists, artist] : artists;
		addArtist.close();
	}

	function removeArtist(id: Artist['id']) {
		const i = artists.findIndex((a) => a.id === id);
		artists = [...artists.slice(0, i), ...artists.slice(i + 1)];
	}
</script>

<Section label="select artists..." grid {hidden} {color}>
	<Option
		label="artists&nbsp;i&nbsp;follow"
		name="followed-artists"
		value="true"
		bind:checked={followedArtists}
		grid
	/>

	{#each artists as artist (artist.id)}
		{#if artist.img}
			<img src={artist.img} alt={artist.name} />
		{/if}
		<input type="checkbox" name="artist" id={artist.id} value={artist.id} checked />
		<label for={artist.id}>{artist.name}</label>
		<button class="delete" type="button" on:click={() => removeArtist(artist.id)}>🗑</button>
	{/each}

	<button class="add" on:click={addArtist.show} type="button">+ add artist</button>
</Section>

<ArtistModal bind:this={addArtist} on:artistSelected={(e) => artistSelected(e.detail.artist)} />

<style>
	img {
		width: 2rem;
		height: 2rem;
		background-color: gray;
		border-radius: 50%;
	}

	input {
		display: none;
	}
	label {
		grid-column: 2;
	}

	.delete {
		background-color: unset;
		border: none;
	}

	.add {
		grid-column: 2;
		padding-top: 0.25rem;
		padding-bottom: 0.25rem;
		padding-right: 1rem;
		padding-left: 1rem;

		color: gray;
		background-color: unset;
		border: 1px dashed gray;
		border-radius: 10rem;
	}
</style>

