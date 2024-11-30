<script lang="ts">
	export let data;
</script>

<svelte:head>
	<title>autofy</title>
	<meta name="description" content="Playlist creator utility for Spotify" />
</svelte:head>

<section>
	<span class="center">No more manual updates!</span>

	<p>
		Playlists created with <span class="autofy">autofy</span> automatically refresh as artists
		release new music.<br /><br />
		<span class="autofy">autofy</span> playlists can also stay up-to-date with your favorite artists,
		adding or removing tracks whenever you follow or unfollow an artist
	</p>

	{#if data.user}
		<a href="/playlists">my&nbsp;playlists</a>
	{:else}
		<form action="/auth/sign-in" method="post" class="ignore">
			<button>connect&nbsp;your&nbsp;Spotify&nbsp;account to&nbsp;get&nbsp;started!</button>
		</form>
	{/if}
</section>

<style>
	section {
		max-width: 35rem;
		margin-right: auto;
		margin-left: auto;
		padding-left: 2rem;
		padding-right: 2rem;
	}

	p {
		margin-top: 0;
		margin-bottom: 4rem;
		text-align: justify;
	}

	.center {
		width: 100%;
		margin-bottom: 3rem;
		text-align: center;
		font-size: 1.2rem;
		color: gray;
		display: block;
	}

	a,
	button {
		width: min-content;
		margin-right: auto;
		margin-left: auto;
		padding: 0.4em 2em;

		color: var(--accent-1);
		background-color: black;
		border: 1px solid var(--accent-1);
		border-radius: 10rem;

		text-align: center;
		line-height: 1.5rem;
		display: block;

		transition:
			color 0.1s ease-in,
			background-color 0.1s ease-in;
	}

	a:hover,
	button:hover {
		color: black;
		background-color: var(--accent-1);
	}

	a {
		font-size: 1.2rem;
	}
</style>
