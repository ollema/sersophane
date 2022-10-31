<script lang="ts">
	import { goto, prefetch } from '$app/navigation';
	import { page } from '$app/stores';

	import { tooltip } from 'svooltip';

	import { formatDate } from '$lib/utils';

	import type { ListResult } from 'pocketbase';
	import type { Event } from '$lib/types';

	export let events: ListResult<Event>;
	export let sort: string | null;

	async function toggleSort(sortBy: string) {
		const url = new URL($page.url);
		const sortParam = url.searchParams.get('sort');
		if (sortParam === `${sortBy}`) {
			url.searchParams.set('sort', `-${sortBy}`);
		} else if (sortParam === `-${sortBy}`) {
			url.searchParams.delete('sort');
		} else {
			url.searchParams.set('sort', `${sortBy}`);
		}
		await prefetch(url.href);
		await goto(url.href);
	}
</script>

<div class="table">
	<table>
		<thead>
			<tr>
				<th
					class:sortAsc={sort === 'name'}
					class:sortDesc={sort === '-name'}
					on:click={() => {
						toggleSort('name');
					}}
				>
					what
				</th>
				<th
					class:sortAsc={sort === 'starts'}
					class:sortDesc={sort === '-starts'}
					on:click={() => {
						toggleSort('starts');
					}}
				>
					when
				</th>
				<th
					class:sortAsc={sort === 'venue'}
					class:sortDesc={sort === '-venue'}
					on:click={() => {
						toggleSort('venue');
					}}
				>
					where
				</th>
				<th>who</th>
				<th />
			</tr>
		</thead>

		<tbody>
			{#each events.items as event}
				<tr>
					<td><a href="/events/{event.id}">{event.name}</a></td>
					<td>{formatDate(event.starts)}</td>
					<td>{event.expand.venue.name}</td>
					<td>
						{#if event.expand['responses(event)'] !== undefined}
							{#each event.expand['responses(event)'] as response}
								<img
									use:tooltip={{ content: response.expand.user.name }}
									src={`http://127.0.0.1:8090/api/files/users/${response.expand.user.id}/${response.expand.user.avatar}?thumb=100x100`}
									alt="avatar for {response.expand.user.name}"
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

<style lang="postcss">
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
		padding: var(--size-2);

		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	th {
		border-bottom: 3px solid var(--surface-2);
	}

	th::after {
		position: relative;
		display: inline-block;
		vertical-align: baseline;
		text-decoration-line: none;

		color: var(--text-2);
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
		color: var(--text-2);
		cursor: pointer;
	}

	tr:nth-child(even) > td {
		background-color: var(--surface-2);
	}

	tr:hover > td {
		background-color: var(--surface-3);
	}

	img {
		position: absolute;
		object-fit: cover;
		height: var(--size-5);
		width: var(--size-5);
	}
</style>
