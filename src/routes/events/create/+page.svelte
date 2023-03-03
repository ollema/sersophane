<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
	$: ({ allVenues, allArtists } = data);
</script>

<svelte:head>
	<title>sersophane • new event</title>
	<meta name="description" content="create new event" />
</svelte:head>

<h3>create new event</h3>

<form action="?/createNewEvent" method="POST">
	<fieldset>
		<label for="name">name</label>
		<input type="text" name="name" id="name" />

		<label for="starts">starts</label>
		<input type="date" name="starts" id="starts" />

		<button type="submit">create</button>
	</fieldset>
</form>

<style lang="postcss">
	h3 {
		margin-bottom: var(--size-3);
	}
	/* form {
		display: flex;
		flex-direction: column;

		& > label {
			margin-top: 1rem;
		}

		& > button {
			margin-top: 1rem;
		}
	} */
</style>
