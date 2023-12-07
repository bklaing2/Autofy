<script lang="ts">
	import GithubLogo from 'virtual:icons/logos/github-icon'
	export let signedIn: boolean
</script>

<footer>
  <ul class="ignore">
		<li><a href="/"><img src="favicon.png" alt="autofy icon" /> autofy</a></li>
		<li><a href="https://github.com/bklaing2/autofy"><GithubLogo style="fill: gray;" /> GitHub</a></li>
		<li><a href="https://buymeacoffee.com/autofy">🍕 support the site!</a></li>
		
		{#if signedIn}
			<li><form action="/auth/sign-out" method="post" class="ignore">
				<button class="sign-out" type="submit">sign out</button>
			</form></li>
		{/if}
  </ul>
</footer>


<style>
	a, li { display: contents; }

	a {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	img {
		width: 1.25rem;
	}


  footer {
		padding: 2rem;
		border-top: 1px solid rgb(30, 30, 30);
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: center;
		align-items: center;
		gap: 2rem;
	}

	footer a {
		font-weight: bold;
	}

	.sign-out { color: gray; }

	/* @media (min-width: 480px) {
		footer {
			padding: 12px 0;
		}
	} */
</style>