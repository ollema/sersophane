<script lang="ts">
	import { goto, prefetch } from '$app/navigation';
	import { page } from '$app/stores';

	import MultiSelect from '$lib/components/MultiSelect.svelte';

	import type { Venue } from '$lib/types';
	import { filtersFromSearchParams } from '$lib/utils';

	let showFilters = false;

	function show() {
		showFilters = !showFilters;
	}

	// filter by name
	let nameFilter = $page.url.searchParams.get('name');

	// filter by venue
	export let allVenues: Venue[];
	const allVenueOptions: { label: string; value: Venue }[] = allVenues.map((venue) => {
		return {
			label: venue.name,
			value: venue
		};
	});
	let selectedVenueOptions: { label: string; value: Venue }[] = filtersFromSearchParams($page.url.searchParams, 'venue', allVenues);

	async function handleFilterSubmit() {
		const url = new URL($page.url);

		if (nameFilter === '' || nameFilter === null) {
			url.searchParams.delete('name');
		} else {
			url.searchParams.set('name', `${nameFilter}`);
		}

		url.searchParams.delete('venue');
		for (const venueFilter of selectedVenueOptions) {
			url.searchParams.append('venue', venueFilter.label as string);
		}

		url.searchParams.delete('page');

		prefetch(url.href);
		await goto(url.href, { keepfocus: true });
	}

	async function clear() {
		const url = new URL($page.url);

		nameFilter = null;
		selectedVenueOptions = [];

		url.searchParams.delete('name');
		url.searchParams.delete('venue');
		url.searchParams.delete('page');

		prefetch(url.href);
		await goto(url.href, { keepfocus: true });
	}
</script>

<div class="filter-component">
	<button on:click={show}>filters</button>

	{#if showFilters}
		<div class="filters">
			<form on:submit|preventDefault={handleFilterSubmit}>
				<div class="inputs">
					<div>
						<label for="name">filter by name:</label>
						<!-- svelte-ignore a11y-autofocus -->
						<input bind:value={nameFilter} id="name" type="text" autocomplete="off" placeholder="event name" autofocus />
					</div>

					<div>
						<label for="venue"> filter by venue: </label>
						<MultiSelect bind:selected={selectedVenueOptions} options={allVenueOptions} placeholder={'venue'} />
					</div>
				</div>

				<div class="buttons">
					<button type="submit">apply filters</button>
					<button type="button" on:click={clear}>clear filters</button>
				</div>
			</form>
		</div>
	{/if}
</div>

<style>
	.filter-component {
		margin-bottom: 1rem;
	}

	button {
		padding: 0.5rem 0.75rem;

		background-color: var(--bg-secondary);
	}

	.filters {
		padding: 1rem;

		border: 3px solid var(--bg-secondary);
	}

	.inputs {
		display: flex;

		flex-wrap: wrap;

		gap: 2rem;

		margin-bottom: 1rem;
	}

	label {
		display: inline-block;
	}

	input {
		display: block;

		max-width: 15rem;

		padding-top: 0.5rem;
		padding-bottom: 0.25rem;

		border-bottom: 3px solid var(--bg-secondary);
	}

	input:hover,
	input:focus {
		border-bottom: 3px solid var(--fg);
	}

	input::placeholder {
		color: var(--text-darker);
	}

	.buttons {
		display: flex;

		flex-wrap: wrap;

		gap: 1rem;
	}
</style>
