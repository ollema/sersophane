<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { slide } from 'svelte/transition';

	let showFilters = false;

	function show() {
		showFilters = !showFilters;
	}

	async function handleSubmit(e: SubmitEvent) {
		if (e.target && e.target instanceof HTMLFormElement) {
			await submitForm(e.target);
		}
	}

	async function submitForm(form: HTMLFormElement) {
		const url = $page.url;

		const name = new FormData(form).get('name');
		const nameParam = name?.toString();
		if (nameParam) {
			url.searchParams.set('name', nameParam);
		} else {
			url.searchParams.delete('name');
		}

		const venue = new FormData(form).get('venue');
		const venueParam = venue?.toString();
		if (venueParam) {
			url.searchParams.set('venue', venueParam);
		} else {
			url.searchParams.delete('venue');
		}

		url.searchParams.delete('page');

		await goto(url.href, { noScroll: true, keepFocus: true });
	}

	let form: HTMLFormElement;

	let nameValue = $page.url.searchParams.get('name');
	let venueValue = $page.url.searchParams.get('venue');
</script>

<div class="filters">
	<button on:click={show}>filters</button>

	{#if showFilters}
		<div transition:slide class="filters-container">
			<form on:submit|preventDefault={handleSubmit} bind:this={form}>
				<div class="filter-inputs">
					<div class="filter-input">
						<label for="name">filter by name</label>
						<input
							id="name"
							name="name"
							on:input={async () => await submitForm(form)}
							type="text"
							autocomplete="event name"
							placeholder="event name"
							value={nameValue}
						/>
					</div>

					<div class="filter-input">
						<label for="venue">filter by venue</label>
						<input
							id="venue"
							name="venue"
							on:input={async () => await submitForm(form)}
							type="text"
							autocomplete="venue name"
							placeholder="venue name"
							value={venueValue}
						/>
					</div>
				</div>
			</form>
		</div>
	{/if}
</div>

<style lang="postcss">
	.filters {
		margin-bottom: var(--size-3);

		& .filters-container {
			padding: var(--size-3);
			margin-top: var(--size-3);

			border: var(--border-size-2) solid var(--surface-3);
			border-radius: var(--radius-2);

			& .filter-inputs {
				display: flex;
				flex-wrap: wrap;
				gap: var(--size-3) var(--size-7);

				& .filter-input {
					display: flex;
					flex-direction: column;
					gap: var(--size-2);

					/* & input {
						padding: var(--size-2);

						background-color: var(--gray-10);
						border: var(--border-size-2) solid var(--surface-3);
						border-radius: var(--radius-2);

						&:focus {
							outline: none;
							background-color: var(--gray-11);
							border-color: var(--brand);
						}
					} */
				}
			}
		}
	}
</style>
