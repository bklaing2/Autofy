<script>
	import Header from './Header.svelte'
	import Footer from './Footer.svelte'
	import './styles.css'
	import { onMount } from 'svelte';

	export let data

	onMount(() => {
		const { hash } = document.location;
		const scrollTo = hash && document.getElementById(hash.slice(1));
		if (scrollTo)
			scrollTo.scrollIntoView();
	});
</script>

<svelte:head>
	<title>autofy</title>
	<meta name="description" content="Playlist creator utility for Spotify" />
</svelte:head>

<div class="layout">
	<Header user={data.user} />
	<slot />
	<Footer signedIn={!!data.user} />
</div>


<style>
	.layout {
		height: 100%;
		max-width: 50rem;
		margin: auto;
		padding-top: 2rem;
		padding-bottom: 2rem;

		display: grid;
		grid-template-columns: 1fr;
		grid-template-rows: min-content minmax(min-content, 1fr) min-content;
		gap: 4rem;
	}
</style>
