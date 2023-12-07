<script lang="ts" context="module">
	interface Playlist {
		id: string;
		title: string;
		artists: { id: string, name: string }[];
		include_followed_artists: boolean;
		update_when_artist_posts: boolean;
		update_when_user_follows_artist: boolean;
		update_when_user_unfollows_artist: boolean;
	};
</script>

<script lang="ts">
	import Section from './Section.svelte';
	import Option from './Option.svelte';
	import Artists, { type Artist } from './Artists.svelte';

  export let action: string
  export let method = 'post'
	
	export let playlist: Playlist = {
		id: '',
		title: '',
		artists: [],
		include_followed_artists: true,
		update_when_artist_posts: true,
		update_when_user_follows_artist: true,
		update_when_user_unfollows_artist: true
	}
	export let color = 'gray'

	let artists: Artist[] = []
	let followedArtists: boolean
</script>


<svelte:head>
	<title>create playlist</title>
	<meta name="description" content="My autofy playlists" />
</svelte:head>


<form {action} {method} class="ignore">
  <input name="title" type="text" placeholder="playlist name" value={playlist.title ?? ''} style:color />

  <Artists bind:artists={artists} bind:followedArtists={followedArtists} {color} />

  <Section label="update when..." {color}>
    <Option label="artist posts new song" name="update-when" value="artist-posts" bind:checked={playlist.update_when_artist_posts} />
    <Option label="i follow new artist" name="update-when" value="user-follows-artist" bind:checked={playlist.update_when_user_follows_artist} disabled={!followedArtists} />
    <Option label="i unfollow artist" name="update-when" value="user-unfollows-artist" bind:checked={playlist.update_when_user_unfollows_artist} disabled={!followedArtists} />
  </Section>

  <button type="submit" class="submit" disabled={artists.length === 0 && !followedArtists}>finish</button>
	<a class="cancel" href="/playlists">cancel</a>
</form>


<style>
	input {
    width: 100%;
		max-width: 30rem;
		padding: 0.5rem;
		font-size: 1.5rem;
		text-align: center;
		color: white;
		border-bottom: 1px solid gray;
		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5)
	}


	.submit {
		margin-top: 4rem;
		padding: 0.5rem;
		padding-left: 2rem;
		padding-right: 2rem;
		
		font-size: 1.2rem;
		color: white;
		background-color: unset;
		border: 1px solid #AB8DF8;
		border-radius: 10rem;
		
		overflow: hidden;
		transition:
			color 0.1s ease-in-out,
			border 0.1s ease-in-out;
	}

  .submit:disabled {
    color: gray;
    border-color: gray;
    cursor: unset;
  }

	.cancel {
		color: rgba(255, 166, 0, 0.7);
		margin-bottom: 4rem;
	}
</style>