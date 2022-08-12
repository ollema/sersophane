<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let showFilters = false;

	export function show() {
		showFilters = !showFilters;
	}

	let nameFilter = $page.url.searchParams.get('name');
	let venueFilter = $page.url.searchParams.get('venue');
	// let profileFilter = $page.url.searchParams.get('profile.name');

	async function handleFilterSubmit() {
		const url = new URL($page.url);

		if (nameFilter === '' || nameFilter === null) {
			url.searchParams.delete('name');
		} else {
			url.searchParams.set('name', `${nameFilter}`);
		}

		if (venueFilter === '' || venueFilter === null) {
			url.searchParams.delete('venue');
		} else {
			url.searchParams.set('venue', `${venueFilter}`);
		}

		url.searchParams.delete('page');

		await goto(url.href);
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
						<input bind:value={nameFilter} id="name" type="text" autocomplete="off" placeholder="event name" />
					</div>

					<div>
						<label for="venue"> filter by venue: </label>
						<input bind:value={venueFilter} id="venue" type="text" autocomplete="off" placeholder="venue" />
					</div>
				</div>

				<button type="submit">apply filters</button>
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

	button:hover {
		text-decoration: underline;
	}

	.filters {
		padding: 1rem;

		border: 2px solid var(--bg-secondary);
	}

	.inputs {
		display: flex;

		flex-wrap: wrap;

		gap: 1.5rem;

		margin-bottom: 1rem;
	}

	label {
		display: inline-block;
	}

	input {
		display: block;

		padding-top: 0.5rem;
		padding-bottom: 0.25rem;

		border-bottom: 2px solid var(--bg-secondary);
	}

	input::placeholder {
		color: var(--text-darker);
	}
</style>
