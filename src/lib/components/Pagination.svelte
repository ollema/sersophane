<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	export let currentPage = 1;
	export let perPage: number;
	export let totalItems: number;

	$: lastPage = Math.max(1, Math.ceil(totalItems / perPage));

	const getRange = (start: number, end: number) => {
		return [...Array(end - start + 1).keys()].map((i) => i + start);
	};

	function getDelta(currentPage: number, lastPage: number) {
		const numberOfButtons = 7;
		let delta: number;
		if (lastPage <= numberOfButtons) {
			delta = numberOfButtons;
		} else {
			delta =
				currentPage > numberOfButtons - 3 && currentPage < lastPage - (numberOfButtons - 4) ? numberOfButtons - 5 : numberOfButtons - 3;
		}
		return delta;
	}

	function pagination(currentPage: number, lastPage: number) {
		if (lastPage === 1) {
			return [1];
		}
		const delta = getDelta(currentPage, lastPage);

		const range = {
			start: Math.round(currentPage - delta / 2),
			end: Math.round(currentPage + delta / 2)
		};

		if (range.start - 1 === 1 || range.end + 1 === lastPage) {
			range.start += 1;
			range.end += 1;
		}

		let pages: number[];

		if (currentPage > delta) {
			pages = getRange(Math.min(range.start, lastPage - delta), Math.min(range.end, lastPage));
		} else {
			pages = getRange(1, Math.min(lastPage, delta + 1));
		}

		if (pages[0] !== 1) {
			if (pages.length + 1 !== lastPage) {
				pages.unshift(-1);
			}
			pages.unshift(1);
		}

		if (pages[pages.length - 1] < lastPage) {
			if (pages.length + 1 !== lastPage) {
				pages.push(-1);
			}
			pages.push(lastPage);
		}

		return pages;
	}

	$: pages = pagination(currentPage, lastPage);

	async function gotoPage(selectedPage: number) {
		const url = new URL($page.url);
		url.searchParams.set('page', `${selectedPage}`);
		await goto(url.href);
	}
</script>

<div class="pagination">
	<button
		type="button"
		on:click={async () => {
			if (currentPage > 1) gotoPage(currentPage - 1);
		}}
	>
		<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
			<path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
		</svg>
	</button>

	{#each pages as _page}
		<button
			type="button"
			class:currentPage={currentPage === _page}
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
			if (currentPage < lastPage) gotoPage(currentPage + 1);
		}}
	>
		<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
			<path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
		</svg>
	</button>
</div>

<style>
	.pagination {
		display: flex;
		align-items: center;
		justify-content: center;

		gap: 0.5rem;

		margin: 0.5rem;
	}

	.pagination button {
		display: flex;
		align-items: center;
		justify-content: center;

		height: 2rem;
		min-width: 2rem;
		padding: 0 0.5rem;

		background-color: var(--bg-secondary);
		font-weight: 600;
	}

	.pagination svg {
		height: 1rem;
		width: 1rem;
	}

	.currentPage {
		color: var(--fg);
	}
</style>
