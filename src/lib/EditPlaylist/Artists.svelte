<script lang="ts" context="module">
	export type Artist = { id: string; name: string; img?: string; };
</script>


<script lang="ts">
  import ArtistModal from '$lib/SearchArtists.svelte'
	import Option from './Option.svelte';

  export let followed = false
  export let artists: Artist[]


  let addArtist: ArtistModal


  function artistSelected (a: SpotifyApi.ArtistObjectFull) {
    const artist: Artist = {
      id: a.id,
      name: a.name,
      img: a.images && a.images.length > 0 ? a.images[0].url : undefined
    }

    artists = !artists.find(a => a.id === artist.id) ? [ ...artists, artist ] : artists
    addArtist.close()
  }

  function removeArtist (id: Artist['id']) {
    const i = artists.findIndex(a => a.id === id)
    artists = [ ...artists.slice(0, i), ...artists.slice(i+1) ]
  }
</script>



<section>
  <h2>artists</h2>
  <Option label="artists&nbsp;i&nbsp;follow" name="followed-artists" value="true" bind:checked={followed} />

  {#each artists as artist (artist.id)}
    {#if artist.img} <img src={artist.img} alt={artist.name} /> {/if}
    <input id={artist.id} name="artist" value={artist.id} type="hidden">
    <span class="name">{artist.name}</span>
    <button class="delete" type="button" on:click={() => removeArtist(artist.id)}>🗑</button>
  {/each}

  <button class="add" on:click={addArtist.show} type="button">+ add artist</button>
</section>

<ArtistModal bind:this={addArtist} on:artistSelected={e => artistSelected(e.detail.artist)} />


<style>
  section {
		width: 100%;
		display: grid;
    grid-template-columns: min-content 1fr min-content;
    gap: 0.5rem;
    align-items: center;
		justify-content: center;

	}


  h2 {
    grid-column: span 3;
    width: 100%;
		color: rgb(200, 200, 200);
		font-size: 1.2rem;
		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
    text-align: center;
	}



  img {
    width: 2rem;
    height: 2rem;
    background-color: gray;
    border-radius: 50%;
  }

  .name {
    grid-column: 2;
    text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
    text-wrap: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .delete { text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5); }

	.add {
		grid-column: span 3;
    justify-self: center;
    width: max-content;
    margin: 5px;
		padding: 0.4em;
		padding-left: 1.6em;
		padding-right: 1.6em;
    font-size: 0.9em;

		color: rgb(255, 255, 255, 0.6);
    text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
		background-color: unset;
		border: 1px dashed rgb(255, 255, 255, 0.6);
		border-radius: 10rem;

    transition: color 0.1s ease-in-out,
    border 0.1s ease-in-out;
	}

  .add:hover {
    color: white;
    border-color: white;
    /* border-style: solid; */
  }
</style>