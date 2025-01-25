<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	let dialog: HTMLDialogElement;
	let shown = false;
	let query = '';
	let artists: SpotifyApi.ArtistObjectFull[] = [];

	export function show() {
		artists = [];
		query = '';
		document.body.classList.add('modal-open');
		dialog.showModal();
		shown = true;
	}

	export function close() {
		dialog.close();
	}

	function onClose() {
		shown = false;
		artists = [];
		query = '';
		document.body.classList.remove('modal-open');
	}

	async function handleKeydown(
		e: KeyboardEvent & { currentTarget: EventTarget & HTMLInputElement }
	) {
		if (e.key === 'Enter') {
			e.preventDefault();
			dispatch('artistSelected', { artist: artists[0] });
		}
		await search();
	}

	async function search() {
		const params = new URLSearchParams({ q: query });
		const response = await fetch(`/api/search-artists?${params}`);
		artists = await response.json();
	}
</script>

<dialog
	bind:this={dialog}
	on:close={onClose}
	on:click={() => dialog.close()}
	role="none"
	class:shown
>
	<div class="layout" on:click={(e) => e.stopPropagation()} role="none">
		<search class="ignore">
			<input
				type="search"
				id="artist"
				placeholder="search..."
				bind:value={query}
				on:keydown={handleKeydown}
			/>
			<button class="search" type="button" on:click={search}>search</button>
		</search>

		{#each artists as artist (artist.id)}
			<button
				class="artist"
				type="button"
				on:click={() => dispatch('artistSelected', { artist: artist })}
			>
				{#if artist.images && artist.images.length > 0}
					<img src={artist.images[0].url} alt={`${artist.name} profile picture`} />
				{/if}
				<span class="name">{artist.name}</span>
			</button>
		{/each}
	</div>
</dialog>

<style>
	dialog {
		height: 100vh;
		width: 100vw;
		max-height: 40rem;
		max-width: 40rem;
		padding: 2rem;

		background-color: unset;
		border: none;
	}
	dialog::backdrop {
		background: linear-gradient(155deg, #ab8df860, #ff00d060);
		backdrop-filter: blur(4px);
		-webkit-backdrop-filter: blur(4px);
		cursor: pointer;
	}

	.layout {
		height: 100%;
		width: 100%;
		padding: 1rem;

		background-color: black;
		border: 6px solid #ab8df8;
		border-style: double;
		border-radius: 12px;

		display: grid;
		grid-template-columns: min-content 1fr min-content;
		grid-auto-rows: min-content;
		align-items: center;
		gap: 1rem;

		overflow: scroll;
	}

	search input {
		grid-column: span 2;
		margin-right: -1rem;
		padding-top: 0.5rem;
		padding-bottom: 0.5rem;
		padding-left: 2rem;
		padding-right: 2rem;
		color: white;
		font-size: 1.2rem;

		border: 1px solid var(--accent-1);
		border-right: none;
		border-top-left-radius: 10rem;
		border-bottom-left-radius: 10rem;
	}

	@property --myColor1 {
		syntax: '<color>';
		initial-value: #ab8df8;
		inherits: false;
	}

	@property --myColor2 {
		syntax: '<color>';
		initial-value: #ab8df8;
		inherits: false;
	}

	search button {
		height: 100%;
		width: 100%;
		padding-left: 1rem;
		padding-right: 1rem;
		color: black;
		background: linear-gradient(155deg, var(--myColor1), var(--myColor2));
		border-top-right-radius: 10rem;
		border-bottom-right-radius: 10rem;
		text-shadow: 0px 0px 0px rgba(0, 0, 0, 0);

		transition:
			color 0.2s ease-in-out,
			text-shadow 0.2s ease-in-out,
			--myColor1 0.2s ease-in-out,
			--myColor2 0.2s ease-in-out;
	}

	search button:hover {
		color: white;
		--myColor1: var(--accent-1);
		--myColor2: var(--accent-2);
		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
	}

	.artist {
		display: contents;
	}

	img {
		height: 2rem;
		aspect-ratio: 1 / 1;
		border-radius: 50%;
	}

	.name {
		grid-column: 2 / span 2;

		text-align: left;
		color: gray;
	}
</style>

