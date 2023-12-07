<script lang="ts">
  export let label: string
  export let hidden = false
  export let className = ''
	export let color = 'gray'
  export let grid = false
</script>



<section class={className} class:grid class:hidden>
  <h2 class="label" style:color>{label}</h2>
  <slot />
</section>


<style>
  section {
		width: 100%;

		display: flex;
		flex-wrap: wrap;
		gap: 1rem;
		justify-content: center;

		overflow: hidden;
	}

  .grid {
		display: grid;
		grid-template-columns: min-content max-content min-content;
		justify-content: center;
    align-items: center;
		gap: 0.75rem;
  }

  .hidden { display: none; }


  h2 {
		width: 100%;
		text-align: center;
		color: gray;
		font-size: 1.2rem;
		font-weight: 100;
		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5)
	}

  .grid h2 { grid-column: span 3; }
</style>