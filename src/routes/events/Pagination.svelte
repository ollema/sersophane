<script lang="ts">
	import { goto, prefetch } from '$app/navigation';
	import { page as pageStore } from '$app/stores';

	import type { ListResult } from 'pocketbase';
	import type { Event } from '$lib/types';

	export let events: ListResult<Event>;

	$: ({ page, totalPages } = events);

	const getRange = (start: number, end: number) => {
		return [...Array(end - start + 1).keys()].map((i) => i + start);
	};

	function getDelta(page: number, totalPages: number) {
		const numberOfButtons = 7;
		let delta: number;
		if (totalPages <= numberOfButtons) {
			delta = numberOfButtons;
		} else {
			delta = page > numberOfButtons - 3 && page < totalPages - (numberOfButtons - 4) ? numberOfButtons - 5 : numberOfButtons - 3;
		}
		return delta;
	}

	function pagination(page: number, totalPages: number) {
		if (totalPages === 1) {
			return [1];
		}
		const delta = getDelta(page, totalPages);

		const range = {
			start: Math.round(page - delta / 2),
			end: Math.round(page + delta / 2)
		};

		if (range.start - 1 === 1 || range.end + 1 === totalPages) {
			range.start += 1;
			range.end += 1;
		}

		let pages: number[];

		if (page > delta) {
			pages = getRange(Math.min(range.start, totalPages - delta), Math.min(range.end, totalPages));
		} else {
			pages = getRange(1, Math.min(totalPages, delta + 1));
		}

		if (pages[0] !== 1) {
			if (pages.length + 1 !== totalPages) {
				pages.unshift(-1);
			}
			pages.unshift(1);
		}

		if (pages[pages.length - 1] < totalPages) {
			if (pages.length + 1 !== totalPages) {
				pages.push(-1);
			}
			pages.push(totalPages);
		}

		return pages;
	}

	$: pages = pagination(page, totalPages);

	async function gotoPage(selectedPage: number) {
		const url = new URL($pageStore.url);
		url.searchParams.set('page', `${selectedPage}`);
		prefetch(url.href);
		await goto(url.href);
	}
</script>

<div class="pagination">
	<button
		type="button"
		on:click={async () => {
			if (page > 1) gotoPage(page - 1);
		}}
	>
		<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
			<path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
		</svg>
	</button>

	{#each pages as _page}
		<button
			type="button"
			class:page={page === _page}
			on:click={async () => {
				if (_page !== -1) {
					gotoPage(_page);
				}
			}}
		>
			{#if _page !== -1}
				{_page}
			{:else}
				...
			{/if}
		</button>
	{/each}

	<button
		type="button"
		on:click={async () => {
			if (page < totalPages) gotoPage(page + 1);
		}}
	>
		<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
			<path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
		</svg>
	</button>
</div>

<style lang="postcss">
	.pagination {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: var(--size-2);

		margin: var(--size-2);

		& button {
			display: flex;
			align-items: center;
			justify-content: center;

			height: var(--size-7);
			min-width: var(--size-7);
			padding: 0 var(--size-2);

			color: var(--link);
			/* background-color: var(--surface-1); */

			&.page {
				color: var(--text-1);
			}
		}
	}

	.pagination svg {
		height: var(--size-3);
		width: var(--size-3);
	}
</style>
