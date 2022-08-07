<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	import type { Venue } from '$lib/types';

	export let venues: Venue[];

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

	$: activeSortParam = $page.url.searchParams.get('sort');

	let nameFilter = $page.url.searchParams.get('name');

	async function handleSubmit() {
		const url = new URL($page.url);
		url.searchParams.set('name', `${nameFilter}`);
		await goto(url.href);
	}
</script>

<form on:submit|preventDefault={handleSubmit}>
	<input bind:value={nameFilter} placeholder="filter by name..." />
</form>

<table>
	<tr>
		<th
			class:sortAsc={activeSortParam === 'name'}
			class:sortDesc={activeSortParam === '-name'}
			on:click={async () => {
				toggleSort('name');
			}}>Name</th
		>
		<th
			class:sortAsc={activeSortParam === 'city'}
			class:sortDesc={activeSortParam === '-city'}
			on:click={async () => {
				toggleSort('city');
			}}>City</th
		>
		<th>Link</th>
	</tr>
	{#each venues as venue}
		<tr>
			<td>{venue.name}</td>
			<td>{venue.city.name}</td>
			<td>{venue.url}</td>
		</tr>
	{/each}
</table>

<style>
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
</style>
