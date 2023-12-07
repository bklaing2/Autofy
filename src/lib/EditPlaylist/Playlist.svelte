<script lang="ts" context="module">
	interface Playlist {
		id: string;
		title: string;
		artists: { id: string, name: string }[];
		followed_artists: boolean;
		update_when_artist_posts: boolean;
		update_when_user_follows_artist: boolean;
		update_when_user_unfollows_artist: boolean;
	};
</script>

<script lang="ts">
	import UpdatesWhen from './UpdatesWhen.svelte'
	import Artists from './Artists.svelte'
	
	export let playlist: Playlist = {
		id: '',
		title: '',
		artists: [],
		followed_artists: true,
		update_when_artist_posts: true,
		update_when_user_follows_artist: true,
		update_when_user_unfollows_artist: true
	}
</script>


<form action={`/playlists/${playlist.id}?/update`} method="post">
  <input name="title" type="text" placeholder="playlist name" value={playlist.title ?? ''} />

	<UpdatesWhen
		bind:artistPosts={playlist.update_when_artist_posts}
		bind:userFollowsArtist={playlist.update_when_user_follows_artist}
		bind:userUnfollowsArtist={playlist.update_when_user_unfollows_artist}
		followedArtists={playlist.followed_artists}
	/>

	<Artists
		bind:followed={playlist.followed_artists}
		bind:artists={playlist.artists}
	/>

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
		border: 2px solid #AB8DF8;
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

	.submit:hover { background-color: #AB8DF8; }
</style>