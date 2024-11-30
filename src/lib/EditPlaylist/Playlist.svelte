<script lang="ts">
	import type { Playlist } from '$lib/types';
	import UpdatesWhen from './UpdatesWhen.svelte';
	import Artists from './Artists.svelte';

	export let playlist: Playlist = {
		id: '',
		title: '',
		artists: [],
		followedArtists: true,
		updateWhenArtistPosts: true,
		updateWhenUserFollowsArtist: true,
		updateWhenUserUnfollowsArtist: true
	};
</script>

<form action={`/playlists/${playlist.id}?/update`} method="post">
	<input name="title" type="text" placeholder="playlist name" value={playlist.title ?? ''} />

	<UpdatesWhen
		bind:artistPosts={playlist.updateWhenArtistPosts}
		bind:userFollowsArtist={playlist.updateWhenUserFollowsArtist}
		bind:userUnfollowsArtist={playlist.updateWhenUserUnfollowsArtist}
		followedArtists={playlist.followedArtists}
	/>

	<Artists bind:followed={playlist.followedArtists} bind:artists={playlist.artists} />

	<button type="submit" class="submit">save</button>
</form>

<style>
	form {
		height: 100%;
		padding: 1rem;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 2rem;

		overflow: scroll;
	}

	input {
		width: 100%;
		max-width: 30rem;
		padding: 0.5rem;
		font-size: 1.5rem;
		text-align: center;
		color: white;
		border-bottom: 1px solid rgb(255, 255, 255, 0.2);
		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
	}

	.submit {
		height: min-content;
		width: 100%;
		margin-top: 2rem;
		margin-bottom: 4rem;
		padding: 0.5rem;
		padding-left: 2rem;
		padding-right: 2rem;

		font-size: 1rem;
		color: white;
		border: 2px solid #ab8df8;
		border-radius: 10rem;

		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
		transition:
			color 0.1s ease-in-out,
			background-color 0.1s ease-in-out,
			border 0.1s ease-in-out;
	}

	.submit:disabled {
		color: gray;
		border-color: gray;
		cursor: unset;
	}

	.submit:hover {
		background-color: #ab8df8;
	}
</style>
