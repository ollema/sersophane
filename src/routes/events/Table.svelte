<script lang="ts">
	import { goto } from '$app/navigation';
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
		await goto(url.href);
	}
</script>

<div class="table">
	<table>
		<thead>
			<tr>
				<th>going?</th>
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
					<td>
						<form method="POST">
							<input type="hidden" name="id" value={event.id} />
							<button type="submit" formaction="?/respondGoing">
								<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
									<path
										fill-rule="evenodd"
										d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
										clip-rule="evenodd"
									/>
								</svg>
							</button>
							<button type="submit" formaction="?/respondInterested">
								<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
									<path
										d="M10.75 4.75a.75.75 0 00-1.5 0v4.5h-4.5a.75.75 0 000 1.5h4.5v4.5a.75.75 0 001.5 0v-4.5h4.5a.75.75 0 000-1.5h-4.5v-4.5z"
									/>
								</svg>
							</button>
							<button type="submit" formaction="?/respondNotGoing">
								<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
									<path
										d="M6.28 5.22a.75.75 0 00-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 101.06 1.06L10 11.06l3.72 3.72a.75.75 0 101.06-1.06L11.06 10l3.72-3.72a.75.75 0 00-1.06-1.06L10 8.94 6.28 5.22z"
									/>
								</svg>
							</button>
						</form>
					</td>
					<td><a href="/events/{event.id}">{event.name}</a></td>
					<td>{formatDate(event.starts)}</td>
					<td>{event.expand.venue.name}</td>
					<td>
						{#if event.expand['responses(event)'] !== undefined}
							{#each event.expand['responses(event)'] as response}
								{#if response.response !== 'not-going'}
									<img
										use:tooltip={{ content: response.expand.user.name }}
										src={`http://127.0.0.1:8090/api/files/users/${response.expand.user.id}/${response.expand.user.avatar}?thumb=100x100`}
										alt="avatar for {response.expand.user.name}"
									/>
								{/if}
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
			minmax(6rem, 6rem)
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

	td > form {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: var(--size-1);

		& > button {
			display: flex;
			align-items: center;
			justify-content: center;

			padding: 0;

			width: var(--size-5);
			height: var(--size-5);

			& > svg {
				width: var(--size-4);
				height: var(--size-4);
			}
		}
	}

	img {
		position: absolute;
		object-fit: cover;
		height: var(--size-5);
		width: var(--size-5);
	}
</style>
