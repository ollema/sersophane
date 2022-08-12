<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	import type { Event, EventResponse } from '$lib/types';
	import { formatToLocalDate } from '$lib/utils';
	import { tooltip } from '$lib/components/tooltip';

	import Pagination from '$lib/components/Pagination.svelte';

	export let events: Event[];
	export let eventResponseMap: { [eventId: string]: EventResponse[] };
	export let currentPage: number | undefined = undefined;
	export let perPage: number;
	export let totalItems: number;
	export let sort: string | null;

	async function toggleSort(sortBy: string) {
		const url = new URL($page.url);
		let sortParam = url.searchParams.get('sort');
		if (sortParam === `${sortBy}`) {
			url.searchParams.set('sort', `-${sortBy}`);
		} else if (sortParam === `-${sortBy}`) {
			url.searchParams.delete('sort');
		} else {
			url.searchParams.set('sort', `${sortBy}`);
		}
		await goto(url.href);
	}
	async function toggleSortByName() {
		toggleSort('name');
	}
	async function toggleSortByStarts() {
		toggleSort('starts');
	}
	async function toggleSortByVenue() {
		toggleSort('venue');
	}

	let nameFilter = $page.url.searchParams.get('name');

	async function handleNameFilterSubmit() {
		const url = new URL($page.url);
		if (nameFilter === '') {
			url.searchParams.delete('name');
		} else {
			url.searchParams.set('name', `${nameFilter}`);
		}
		url.searchParams.delete('page');
		await goto(url.href);
	}
</script>

<main>
	<form on:submit|preventDefault={handleNameFilterSubmit}>
		<input bind:value={nameFilter} placeholder="filter by event name..." />
	</form>

	<table>
		<thead>
			<tr>
				<th on:click={toggleSortByName}>what</th>
				<th on:click={toggleSortByStarts}>when</th>
				<th on:click={toggleSortByVenue}>where</th>
				<th>who</th>
				<th />
			</tr>
		</thead>

		<tbody>
			{#each events as event}
				<tr>
					<td><a href="/events/{event.id}">{event.name}</a></td>
					<td>{formatToLocalDate(event.starts, 'd/M')}</td>
					<td>{event.venue.name}</td>
					<td>
						{#if eventResponseMap[event.id] !== undefined}
							{#each eventResponseMap[event.id] as response}
								<img use:tooltip title={response.profile.name} src={response.profile.avatar} alt="avatar for {response.profile.name}" />
							{/each}
						{/if}
					</td>
					<td />
				</tr>
			{/each}
		</tbody>
	</table>
</main>

<footer>
	<Pagination {currentPage} {perPage} {totalItems} />
</footer>

<style>
	input {
		padding: 0.5rem;
		margin-bottom: 1rem;

		background-color: var(--bg-secondary);
	}

	table {
		display: grid;

		grid-template-columns:
			minmax(2rem, 2fr)
			minmax(2rem, 1fr)
			minmax(2rem, 2fr)
			minmax(2rem, 3fr)
			minmax(2rem, 1fr);

		min-width: 40rem;
	}

	thead,
	tbody,
	tr {
		display: contents;
	}

	td,
	th {
		position: relative;

		display: flex;
		align-items: center;

		padding: 0.5rem;

		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	th {
		color: var(--fg);
		font-weight: 600;
		border-bottom: 2px solid var(--bg-secondary);
	}

	th:hover {
		color: var(--fg-hover);
	}

	tr:nth-child(even) > td {
		background-color: var(--bg-secondary);
	}

	tr:hover > td {
		background-color: var(--bg-hover);
	}

	img {
		position: absolute;

		object-fit: cover;
		height: 1.5rem;
		width: 1.5rem;
	}
</style>
