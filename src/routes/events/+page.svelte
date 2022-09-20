<script lang="ts">
	import Filters from './Filters.svelte';

	import { goto, prefetch } from '$app/navigation';
	import { page } from '$app/stores';

	import { tooltip } from 'svooltip';

	import { formatToLocalDate } from '$lib/utils';
	import Pagination from '$lib/components/Pagination.svelte';

	import type { PageData } from './$types';

	export let data: PageData;
	$: ({ events, eventResponseMap, allVenues, currentPage, perPage, totalItems, sortBy } = data);

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
		prefetch(url.href);
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

<svelte:head>
	<title>sersophane • events</title>
</svelte:head>

<main>
	<Filters {allVenues} />

	<div class="table">
		<table>
			<thead>
				<tr>
					<th class:sortAsc={sortBy === 'name'} class:sortDesc={sortBy === '-name'} on:click={toggleSortByName}>what</th>
					<th class:sortAsc={sortBy === 'starts'} class:sortDesc={sortBy === '-starts'} on:click={toggleSortByStarts}>when</th>
					<th class:sortAsc={sortBy === 'venue'} class:sortDesc={sortBy === '-venue'} on:click={toggleSortByVenue}>where</th>
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
									<img
										use:tooltip={{ content: response.profile.name }}
										src={response.profile.avatar}
										alt="avatar for {response.profile.name}"
									/>
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

<style lang="postcss">
	.table {
		@apply overflow-auto;
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
		@apply contents;
	}

	td,
	th {
		@apply relative flex items-center p-2 overflow-hidden text-ellipsis whitespace-nowrap;
	}

	th {
		@apply text-zinc-500 font-semibold border-b-2 border-zinc-800 select-none;
	}

	th::after {
		@apply inline-block relative text-xs align-baseline no-underline text-zinc-500;
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
		@apply hover:text-zinc-300 cursor-pointer;
	}

	tr:nth-child(even) > td {
		@apply bg-zinc-800;
	}

	tr:hover > td {
		@apply bg-zinc-700;
	}

	img {
		@apply absolute object-cover h-6 w-6;
		position: absolute;
	}
</style>
