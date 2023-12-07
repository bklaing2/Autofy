<script labg="ts">
	import Playlists from "./Playlists.svelte"
	export let data
	let playlists = data.playlists ?? []
</script>

<svelte:head>
	<script data-name="BMC-Widget" data-cfasync="false" src="https://cdnjs.buymeacoffee.com/1.0.0/widget.prod.min.js" data-id="autofy" data-description="Support me on Buy me a coffee!" data-color="#AB8DF8" data-position="Right" data-x_margin="18" data-y_margin="18"></script>
</svelte:head>


<Playlists playlists={playlists}>
	<slot />
</Playlists>