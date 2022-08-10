<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	import type { Event } from '$lib/types';
	import { formatToLocalDate } from '$lib/utils';

	import Pagination from '$lib/components/Pagination.svelte';

	export let events: Event[];
	export let currentPage: number;
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
		<input bind:value={nameFilter} placeholder="filter by name..." />
	</form>

	<table>
		<tr>
			<th>Date</th>
			<th
				class:sortAsc={sort === 'name'}
				class:sortDesc={sort === '-name'}
				on:click={async () => {
					toggleSort('name');
				}}>Name</th
			>
			<th
				class:sortAsc={sort === 'venue'}
				class:sortDesc={sort === '-venue'}
				on:click={async () => {
					toggleSort('venue');
				}}>Venue</th
			>
			<th>Artists</th>
			<th>Type</th>
			<th>Link</th>
		</tr>
		{#each events as event}
			<tr>
				<td>{formatToLocalDate(event.startsAt, 'd/M yyyy')}</td>
				<td>{event.name}</td>
				<td>{event.venue.name}</td>
				<td>
					{#each event.artists as artist, i}
						{artist.name}{#if i < event.artists.length - 2},&nbsp;{/if}{#if i === event.artists.length - 2}&nbsp;&&nbsp;{/if}
					{/each}
				</td>
				<td>{event.type}</td>
				<td>{event.url}</td>
			</tr>
		{/each}
	</table>
</main>

<footer>
	<Pagination {currentPage} {perPage} {totalItems} />
</footer>

<style>
	input {
		padding: 0.5rem;
		margin-bottom: 1rem;

		background-color: var(--background-secondary);
	}

	th {
		padding: 0.5rem;

		background-color: var(--background-secondary);
	}

	th:hover {
		text-decoration: underline;
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
		top: 0.1em;
		text-decoration: none;
	}

	th.sortDesc::after {
		content: '\a0\a0▲';
		top: -0.4em;
		text-decoration: none;
	}

	td {
		padding: 0.5rem;
	}

	tr:hover {
		background-color: var(--background-tertiary);
	}
</style>
