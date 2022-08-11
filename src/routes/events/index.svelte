<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	import type { Event, EventResponse } from '$lib/types';
	import { formatToLocalDate } from '$lib/utils';

	import Pagination from '$lib/components/Pagination.svelte';

	export let events: Event[];
	export let responses: { [eventId: string]: EventResponse[] };
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
		<tr>
			<th
				class:sortAsc={sort === 'name'}
				class:sortDesc={sort === '-name'}
				class="sortable"
				on:click={async () => {
					toggleSort('name');
				}}>what</th
			>
			<th
				class:sortAsc={sort === 'starts'}
				class:sortDesc={sort === '-starts'}
				class="sortable"
				on:click={async () => {
					toggleSort('starts');
				}}>when</th
			>
			<th
				class:sortAsc={sort === 'venue'}
				class:sortDesc={sort === '-venue'}
				class="sortable"
				on:click={async () => {
					toggleSort('venue');
				}}>where</th
			>
			<th>who</th>
		</tr>
		{#each events as event}
			<tr>
				<td><a href="/events/{event.id}">{event.name}</a></td>
				<td>{formatToLocalDate(event.starts, 'd/M yyyy')}</td>
				<td>{event.venue.name}</td>
				<td class="avatars">
					{#if responses !== undefined && responses[event.id] !== undefined}
						{#each responses[event.id] as response}
							<img src={response.profile.avatar} alt="avatar for {response.profile.name}" />
						{/each}
					{/if}
				</td>
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

	th.sortable:hover {
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

	td a:hover {
		text-decoration: underline;
	}

	tr:hover {
		background-color: var(--background-tertiary);
	}

	.avatars {
		position: relative;

		padding: 0rem;
	}

	img {
		position: absolute;
		object-fit: cover;
		height: 2rem;
		width: 2rem;
	}
</style>
