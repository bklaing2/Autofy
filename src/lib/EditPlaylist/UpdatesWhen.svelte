<script lang="ts">
	import Option from './Option.svelte';

	export let artistPosts: boolean;
	export let userFollowsArtist: boolean;
	export let userUnfollowsArtist: boolean;

	export let followedArtists: boolean;
</script>

<section>
	<h2>updates when</h2>
	<Option
		label="artist releases music"
		name="update-when"
		value="artist-posts"
		bind:checked={artistPosts}
	/>
	<Option
		label="i follow an artist"
		name="update-when"
		value="user-follows-artist"
		bind:checked={userFollowsArtist}
		disabled={!followedArtists}
	/>
	<Option
		label="i unfollow an artist"
		name="update-when"
		value="user-unfollows-artist"
		bind:checked={userUnfollowsArtist}
		disabled={!followedArtists}
	/>
</section>

<style>
	section {
		width: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	h2 {
		width: 100%;
		margin-bottom: 0.5rem;
		color: rgb(200, 200, 200);
		font-size: 1.2rem;
		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
		text-align: center;
	}
</style>

