<script lang="ts">
	export let label: string;
	export let name: string;
	export let value: string;
	export let checked: boolean;
	export let disabled = false;
</script>

<input {name} id={value} {value} type="checkbox" bind:checked {disabled} />
<label for={value} class:disabled>{label}</label>

<style>
	input {
		display: none;
	}

	label {
		grid-column: span 3;
		margin: 5px;
		padding: 0.6em;
		padding-left: 1.6em;
		padding-right: 1.6em;
		text-align: center;

		font-size: 0.9rem;
		background-color: black;
		border-radius: 10em;
		position: relative;

		text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
		box-shadow: 0px 0px 0px 3px rgba(0, 0, 0, 0);
		-webkit-tap-highlight-color: rgba(0, 0, 0, 0);
		user-select: none;
		cursor: pointer;

		transition:
			background-color 0.1s ease-in,
			text-shadow 0.1s ease-in,
			box-shadow 0.1s ease-in;
	}

	label.disabled {
		color: gray;
		cursor: unset;
	}

	label::before {
		content: '';
		position: absolute;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		z-index: -1;
		margin: -2px;
		border-radius: inherit;
		background: linear-gradient(155deg, #ab8df8, #ff00d0);
	}

	label.disabled::before {
		background: unset;
		background-color: gray;
	}

	@media only screen and (min-width: 30rem) {
		label:not(.disabled):hover {
			background-color: unset;
			text-shadow: 1px 0px 1px rgba(0, 0, 0, 0.5);
		}
	}

	input:checked + label:not(.disabled) {
		background-color: unset;
		box-shadow:
			5px 5px rgba(0, 0, 0, 0.2),
			inset 1px 0.5px 2px white,
			inset -1px -1px 3px black;
	}
</style>

