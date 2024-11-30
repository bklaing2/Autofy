<script lang="ts">
	export let id: string;
	export let flipped = false;
</script>

<div class="playlist" class:flipped>
	<div class="face">
		<iframe
			title="playlist"
			src={`https://open.spotify.com/embed/playlist/${id}`}
			height="100%"
			frameBorder="0"
			allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
			loading="lazy"
		/>

		{#if !flipped}
			<a class="flip hidden" href={`/playlists/${id}`} data-sveltekit-noscroll>⚙&#xFE0E;</a>
		{/if}
	</div>

	<div class="settings face">
		<slot />
		<a class="flip" href="/playlists" data-sveltekit-noscroll>↩&#xFE0E;</a>
	</div>
</div>

<style>
	.playlist {
		width: 300px;
		max-width: 100%;
		max-height: 100%;
		aspect-ratio: 1 / 1.414;

		background: linear-gradient(155deg, #ab8df8, #ff00d0);
		border-radius: 12px;

		position: relative;
		transform-style: preserve-3d;
		-webkit-transform-style: preserve-3d;
		transition: all 0.5s linear;
		-webkit-transition: all 0.5s linear;
	}

	.face {
		position: absolute;
		width: 100%;
		height: 100%;

		border-radius: inherit;
		backface-visibility: hidden;
		-webkit-backface-visibility: hidden;
	}

	.settings {
		overflow: hidden;
	}

	.flipped,
	.settings {
		transform: rotateY(180deg);
		-webkit-transform: rotateY(180deg);
	}

	.flip {
		position: absolute;
		bottom: 5px;
		right: 5px;
		width: 2.5rem;
		height: 2.5rem;

		font-size: 1.5rem;
		line-height: 1.5rem;
		color: white;
		background-color: unset;
		border: none;
		border-radius: 50%;

		display: flex;
		justify-content: center;
		align-items: center;

		backdrop-filter: blur(2px);
		-webkit-backdrop-filter: blur(2px);
		background-color: #ab8df850;
		text-shadow:
			-0.1px -0.1px 0px white,
			-0.1px -0.1px 1px black;
		box-shadow:
			inset 1.5px 1.5px 3px 0px white,
			inset -2px -2px 3px 0px black;

		opacity: 1;
		cursor: pointer;
		transition:
			color 0.2s ease-in-out,
			opacity 0.2s ease-in-out;
	}

	.hidden {
		opacity: 0;
	}

	.playlist:hover .flip {
		opacity: 1;
	}
	.playlist:hover .flip:hover {
		opacity: 0.7;
	}

	@media only screen and (max-width: 30rem) {
		.flip {
			opacity: 1;
		}
	}
</style>

