<script lang="ts">
	import MultiSelect from 'svelte-multiselect';
	import type { Option } from 'svelte-multiselect';

	import type { PageData } from './$types';

	export let data: PageData;
	$: ({ allVenues, allArtists } = data);

	let venue: Option[];
	let artists: Option[];
</script>

<svelte:head>
	<title>sersophane • new event</title>
	<meta name="description" content="create new event" />
</svelte:head>

<h2>create new event</h2>

<form action="?/createNewEvent" method="POST">
	<label for="name">name</label>
	<input type="text" name="name" id="name" />

	<label for="starts">starts</label>
	<input type="date" name="starts" id="starts" />

	<label for="venue">venue</label>
	<MultiSelect bind:selected={venue} options={allVenues} />

	<label for="artists">artists</label>
	<MultiSelect bind:selected={artists} options={allArtists} />

	<button type="submit">create</button>
</form>

<style lang="postcss">
	form {
		display: flex;
		flex-direction: column;

		& > label {
			margin-top: 1rem;
		}

		& > button {
			margin-top: 1rem;
		}
	}
</style>
