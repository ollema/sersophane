<script lang="ts">
	import { invalidateAll } from '$app/navigation';

	import type { Profile } from '$lib/types';

	export let section: string;
	export let profile: Profile | undefined;

	async function signout() {
		await fetch('auth/signout', { method: 'POST' });
		console.log('signing out - await invalidateAll()');
		await invalidateAll();
		console.log('invalidateAll() awaited!');
	}
</script>

<nav>
	<div>
		<div class="left-nav">
			<a data-sveltekit-prefetch href="/" class:selected={section === ''}>sersophane</a>
			<a data-sveltekit-prefetch href="/events" class:selected={section === 'events'}>events</a>
			<a data-sveltekit-prefetch href="/about" class:selected={section === 'about'}>about</a>
		</div>

		<div class="right-nav">
			{#if profile}
				<a href="/">{profile.name}</a>
				<button on:click={signout} type="button">sign out</button>
			{:else}
				<a href="/signin">sign in</a>
			{/if}
		</div>
	</div>
</nav>

<style>
	nav {
		color: var(--text);
		font-weight: 600;
		border-bottom: 3px solid var(--bg-secondary);
	}

	nav > div {
		display: flex;
		align-items: center;
		justify-content: space-between;

		max-width: 1080px;
		margin: 0 auto;
	}

	.left-nav {
		display: flex;
		align-items: center;

		gap: 1rem;
	}

	.right-nav {
		display: flex;
		align-items: center;

		gap: 1rem;
	}

	a,
	button {
		padding: 1rem 0.5rem 0.75rem 0.5rem;

		border-bottom: 4px solid rgba(0, 0, 0, 0);
		transition: border-color 0.2s ease-in-out, color 0.2s ease-in-out;
	}

	a.selected {
		color: var(--fg);
		border-color: var(--fg);
	}

	a.selected:hover,
	a.selected:focus {
		color: var(--fg-hover);
		border-color: var(--fg-hover);
	}
</style>
