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

<style lang="postcss">
	.filter-component {
		@apply mb-4;
	}

	button {
		@apply py-2 px-3 bg-zinc-800;
	}

	.filters {
		@apply p-4 border-2 border-zinc-800;
	}

	.inputs {
		@apply flex flex-wrap gap-8 mb-4;
	}

	label {
		@apply inline-block;
	}

	input {
		@apply block w-full max-w-[15rem] bg-zinc-900 border-0 border-b-2 border-zinc-800 focus:ring-0 focus:border-emerald-600 hover:border-emerald-300;
	}

	input:-webkit-autofill,
	input:-webkit-autofill:hover,
	input:-webkit-autofill:focus {
		-webkit-text-fill-color: #a1a1aa;
		caret-color: #a1a1aa;
		box-shadow: 0 0 0px 1000px #18181b inset;
		-webkit-box-shadow: 0 0 0px 1000px #18181b inset;
	}

	input::placeholder {
		@apply text-zinc-500;
	}

	.buttons {
		@apply flex flex-wrap gap-4;
	}
</style>
