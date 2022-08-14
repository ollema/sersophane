<script lang="ts">
	import { tooltip } from '$lib/components/tooltip';

	export let searchText = '';
	export let showOptions = false;
	export let options: Option[];
	export let matchingOptions: Option[] = [];

	export let selected: Option[] = [];
	export let selectedLabels: (string | number)[] = [];
	export let selectedValues: unknown[] = [];

	export let input: HTMLInputElement | null = null;
	export let outerDiv: HTMLDivElement | null = null;
	export let placeholder: string | undefined = undefined;
	export let activeOption: Option | null = null;
	export let filterFunc = (op: Option, searchText: string) => {
		if (!searchText) return true;
		return `${op.label}`.toLowerCase().includes(searchText.toLowerCase());
	};

	export let autoScroll = true;
	export let required = false;
	export let autocomplete = `off`;
	export let invalid = false;

	type Option = {
		label: string;
		value: unknown;
	};

	$: selectedLabels = selected.map((op) => op.label);
	$: selectedValues = selected.map((op) => op.value);

	// formValue binds to input.form-control to prevent form submission if required is true and no options are selected
	$: formValue = selectedValues.join(`,`);
	$: if (formValue) invalid = false; // reset error status whenever component state changes

	// options matching the current search text
	$: matchingOptions = options.filter(
		// remove already selected options from dropdown list
		(op) => filterFunc(op, searchText) && !selectedLabels.includes(op.label)
	);

	// reset activeOption if it's no longer in the matchingOptions list
	$: if (activeOption && !matchingOptions.includes(activeOption)) activeOption = null;

	// add an option to selected list
	function add(option: Option) {
		searchText = ``;
		selected = [...selected, option];
		input?.focus();
	}

	// remove an option from selected list
	function remove(option: Option) {
		if (selected.length === 0) return;
		selected.splice(selectedLabels.lastIndexOf(option.label), 1);
		selected = selected; // Svelte rerender
	}

	function setOptionsVisible(show: boolean) {
		showOptions = show;
		if (show) {
			input?.focus();
		} else {
			input?.blur();
			activeOption = null;
		}
	}

	// handle all keyboard events this component receives
	async function handleKeydown(event: KeyboardEvent) {
		// on escape or tab out of input: dismiss options dropdown and reset search text
		if (event.key === `Escape` || event.key === `Tab`) {
			setOptionsVisible(false);
			searchText = ``;
		}

		// on enter key: toggle active option and reset search text
		else if (event.key === `Enter`) {
			// prevent enter key from triggering form submission
			event.preventDefault();

			if (activeOption) {
				selectedLabels.includes(activeOption.label) ? remove(activeOption) : add(activeOption);
				searchText = ``;
			}

			// no active option & no search text -> options dropdown is closed -> enter should open it
			else setOptionsVisible(true);
		}

		// on up/down arrow keys: update active option
		else if ([`ArrowDown`, `ArrowUp`].includes(event.key)) {
			// if no option is active yet, but there are matching options, make first one active
			if (activeOption === null && matchingOptions.length > 0) {
				activeOption = matchingOptions[0];
				return;
			} else if (activeOption === null) {
				// if no option is active and no options are matching, do nothing
				return;
			}
			const increment = event.key === `ArrowUp` ? -1 : 1;
			const newActiveIdx = matchingOptions.indexOf(activeOption) + increment;

			if (newActiveIdx < 0) {
				// wrap around top
				activeOption = matchingOptions[matchingOptions.length - 1];
			} else if (newActiveIdx === matchingOptions.length) {
				// wrap around bottom
				activeOption = matchingOptions[0];
			} else {
				// default case: select next/previous in item list
				activeOption = matchingOptions[newActiveIdx];
			}
			if (autoScroll) {
				// TODO This ugly timeout hack is needed to properly scroll element into view when wrapping
				// around start/end of option list. Find a better solution than waiting 10 ms to.
				setTimeout(() => {
					const li = document.querySelector(`ul.options > li.active`);
					if (li) {
						// eslint-disable-next-line @typescript-eslint/ban-ts-comment
						// @ts-ignore
						li.parentNode?.scrollIntoView({ block: `center` });
						// eslint-disable-next-line @typescript-eslint/ban-ts-comment
						// @ts-ignore
						li.scrollIntoViewIfNeeded();
					}
				}, 10);
			}
		}

		// on backspace key: remove last selected option
		else if (event.key === `Backspace` && selected.length > 0 && !searchText) {
			remove(selected.at(-1) as Option);
		}
	}

	$: is_selected = (label: string | number) => selectedLabels.includes(label);

	const if_enter_or_space = (handler: () => void) => (event: KeyboardEvent) => {
		if ([`Enter`, `Space`].includes(event.code)) {
			event.preventDefault();
			handler();
		}
	};
</script>

<svelte:window
	on:click={(event) => {
		// eslint-disable-next-line @typescript-eslint/ban-ts-comment
		// @ts-ignore
		if (outerDiv && !outerDiv.contains(event.target)) {
			setOptionsVisible(false);
		}
	}}
/>

<div
	bind:this={outerDiv}
	class:open={showOptions}
	aria-expanded={showOptions}
	aria-multiselectable
	class:invalid
	class="multiselect"
	on:mouseup|stopPropagation={() => setOptionsVisible(true)}
>
	<input
		{required}
		bind:value={formValue}
		tabindex="-1"
		aria-hidden="true"
		aria-label="ignore this, used only to prevent form submission if select is required but empty"
		class="form-control"
		on:invalid={() => (invalid = true)}
	/>

	<ul class="selected">
		{#each selected as option}
			<li aria-selected="true">
				{option.label}
				<button
					use:tooltip
					title="remove {option.label}"
					class="remove-icon"
					on:mouseup|stopPropagation={() => remove(option)}
					on:keydown={if_enter_or_space(() => remove(option))}
					type="button"
				>
					<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 22.5 22.5" stroke="currentColor" stroke-width="2">
						<path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
					</svg>
				</button>
			</li>
		{/each}
		<li style="display: contents;">
			<input
				bind:this={input}
				{autocomplete}
				bind:value={searchText}
				on:mouseup|self|stopPropagation={() => setOptionsVisible(true)}
				on:keydown={handleKeydown}
				on:focus={() => setOptionsVisible(true)}
				placeholder={selectedLabels.length ? `` : placeholder}
				aria-invalid={invalid ? `true` : null}
			/>
		</li>
	</ul>

	{#if searchText || options?.length > 0}
		<ul class:hidden={!showOptions} class="options">
			{#each matchingOptions as option}
				<li
					on:mousedown|stopPropagation
					on:mouseup|stopPropagation={() => {
						is_selected(option.label) ? remove(option) : add(option);
					}}
					class:selected={is_selected(option.label)}
					class:active={activeOption && activeOption.label === option.label}
					on:mouseover={() => {
						activeOption = option;
					}}
					on:focus={() => {
						activeOption = option;
					}}
					on:mouseout={() => (activeOption = null)}
					on:blur={() => (activeOption = null)}
					aria-selected="false"
				>
					{option.label}
				</li>
			{/each}
		</ul>
	{/if}
</div>

<style>
	div.multiselect {
		position: relative;

		align-items: center;
		display: flex;

		max-width: 15rem;

		padding: 0.5rem 0 0.25rem 0;

		cursor: text;
		border-bottom: 3px solid var(--bg-secondary);
		background-color: var(--bg);
	}

	div.multiselect:hover {
		border-bottom: 3px solid var(--fg);
	}

	div.multiselect.open {
		z-index: 4;
	}

	div.multiselect > ul.selected {
		display: flex;

		flex-wrap: wrap;

		gap: 0.5rem;
	}

	div.multiselect > ul.selected > li {
		display: flex;
		align-items: center;

		padding: 0 0.25rem 0 0.25rem;

		white-space: nowrap;

		background-color: var(--bg-secondary);
	}

	div.multiselect > ul.selected > li > button.remove-icon {
		display: flex;
		align-items: center;
	}

	div.multiselect > ul.selected > li > button.remove-icon > svg {
		height: 1rem;
		width: 1rem;

		color: var(--text);
	}

	div.multiselect > ul.selected > li > button.remove-icon > svg:hover,
	div.multiselect > ul.selected > li > button.remove-icon:focus > svg {
		color: red;
	}

	div.multiselect > ul.selected > li > input {
		flex: 1;
		min-width: 1rem;
	}

	div.multiselect > ul.selected > li > input::placeholder {
		color: var(--text-darker);
	}

	div.multiselect > input.form-control {
		position: absolute;
		background: transparent;

		border: none;
		outline: none;
		z-index: -1;
		opacity: 0;
		pointer-events: none;
	}

	div.multiselect > ul.options {
		position: absolute;
		top: 100%;
		left: 0;

		width: 100%;
		max-height: 50vh;

		overflow: auto;
		background-color: var(--bg-secondary);
		color: var(--text);
	}

	div.multiselect > ul.options.hidden {
		visibility: hidden;
		opacity: 0;
	}

	div.multiselect > ul.options > li {
		padding: 0.5rem 0;
		cursor: pointer;
	}

	div.multiselect > ul.options > li.active {
		background: var(--bg-hover);
	}
</style>
