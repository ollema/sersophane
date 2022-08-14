<script lang="ts">
	import Filters from './_Filters.svelte';

	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	import type { Event, EventResponse, Venue } from '$lib/types';
	import { formatToLocalDate } from '$lib/utils';
	import { tooltip } from '$lib/components/tooltip';

	import Pagination from '$lib/components/Pagination.svelte';

	export let events: Event[];
	export let eventResponseMap: { [eventId: string]: EventResponse[] };
	export let allVenues: Venue[];
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
</script>

<main>
	<Filters {allVenues} />

	<div class="table">
		<table>
			<thead>
				<tr>
					<th class:sortAsc={sort === 'name'} class:sortDesc={sort === '-name'} on:click={toggleSortByName}>what</th>
					<th class:sortAsc={sort === 'starts'} class:sortDesc={sort === '-starts'} on:click={toggleSortByStarts}>when</th>
					<th class:sortAsc={sort === 'venue'} class:sortDesc={sort === '-venue'} on:click={toggleSortByVenue}>where</th>
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
	</div>
</main>

<footer>
	<Pagination {currentPage} {perPage} {totalItems} />
</footer>

<style>
	.table {
		overflow: auto;
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
		border-bottom: 3px solid var(--bg-secondary);
	}

	th::after {
		display: inline-block;
		font-size: xx-small;
		vertical-align: baseline;
		position: relative;
		text-decoration: none;
		color: var(--slate-400);
	}

	th.sortAsc::after {
		content: '\a0\a0▼';
		top: 0.4em;
		text-decoration: none;
	}

	th.sortDesc::after {
		content: '\a0\a0▲';
		top: -0.4em;
		text-decoration: none;
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
