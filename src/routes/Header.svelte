<script lang="ts">
	export let user: SpotifyApi.CurrentUsersProfileResponse | undefined
	const userImage = user?.images && user.images.length > 0 ? user.images[0].url : undefined

	const href = user ? "/playlists" : "/"
</script>

<header>
	<a {href}><img src="favicon.png" alt="autofy icon" /></a>
	<a class="site-title autofy" {href}>autofy</a>

	{#if user}
		{#if userImage} <img class="profile" src={userImage} alt={`${user.display_name} profile picture`} />
		{:else} <span class="profile secondary">{user.display_name}</span> {/if}

		<form action="/auth/sign-out" method="post" class="ignore">
			<button class="sign-out" type="submit">sign out</button>
		</form>
	{/if}
</header>

<style>
	header {
		position: relative;
		padding-bottom: 2rem;
		padding-left: 2rem;
		padding-right: 2rem;

		display: grid;
		grid-template-columns: repeat(2, max-content) 1fr;
		justify-items: end;
		align-items: center;
		gap: 2rem;
		overflow: scroll;
	}


	img {
		height: 0.5rem;
		aspect-ratio: 1 / 1;
	}
	.site-title {
		font-weight: 900;
		font-size: 2rem;
	}

	.secondary { color: gray; }

	img {
		height: 2rem;
		aspect-ratio: 1 / 1;
		border-radius: 50%;
	}


	.sign-out {
		position: absolute;
		bottom: 0;
		right: 1.5rem;
		padding: 0.25rem;
		padding-left: 1rem;
		padding-right: 1rem;

		font-size: 0.8rem;
		color: white;
		background-color: gray;
		border-radius: 10rem;
		visibility: hidden;

		transition: visibility 0.2s ease-out;
	}

	.sign-out::after {
		position: absolute;
		top: -0.6rem;
		right: 1rem;
    content: '▲';
		color: gray;
	}

	.profile:hover + form .sign-out, .sign-out:hover { visibility: unset; }
</style>
