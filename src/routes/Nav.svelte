<script lang="ts">
	import { invalidateAll } from '$app/navigation';

	import type { Profile } from '$lib/types';

	export let section: string;
	export let profile: Profile | undefined;

	async function signout() {
		await fetch('auth/signout', { method: 'POST' });
		await invalidateAll();
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
				<a class:selected={section === 'signin'} href="/signin">sign in</a>
			{/if}
		</div>
	</div>
</nav>

<style lang="postcss">
	nav {
		@apply text-zinc-300 font-semibold border-b-2 border-zinc-800;
	}

	nav > div {
		@apply flex items-center justify-between max-w-6xl my-0 mx-auto;
	}

	.left-nav,
	.right-nav {
		@apply flex items-start justify-center gap-4;
	}

	a,
	button {
		@apply px-2 pt-4 pb-3 border-b-4 border-transparent transition-colors hover:text-emerald-400;
	}

	a.selected {
		@apply text-emerald-600 border-emerald-600 hover:text-emerald-400 hover:border-emerald-400;
	}
</style>
