<script lang="ts">
	import { page } from '$app/stores';

	import 'svooltip/styles.css';
	import '../openprops.css';
	import '../app.css';

	import type { LayoutServerData } from './$types';

	export let data: LayoutServerData;
	$: ({ user } = data);

	$: section = $page.url.pathname.split('/')[1];
</script>

<nav>
	<div>
		<ul class="links">
			<li><a href="/" class:selected={section === ''}>sersophane</a></li>
			<li><a href="/events" class:selected={section === 'events'}>events</a></li>
			<li><a href="/about" class:selected={section === 'about'}>about</a></li>
		</ul>

		<ul class="profile">
			{#if user}
				<li><a href="/account">{user?.name || user?.username || 'no username set'}</a></li>
			{:else}
				<li><a href="/auth/signin">sign in</a></li>
			{/if}
		</ul>
	</div>
</nav>

<main>
	<slot />
</main>

<style lang="postcss">
	nav {
		max-width: var(--size-lg);
		margin: 0 auto;
		padding: var(--size-2) 0;

		& div {
			display: flex;
			align-items: center;
			justify-content: space-between;
		}
	}

	ul {
		display: flex;
		align-items: center;

		padding: 0;

		list-style: none;
	}

	li {
		padding: var(--size-2) var(--size-3);
	}

	a {
		color: var(--link);
		font-size: var(--font-size-3);
		font-weight: var(--font-weight-6);

		&.selected {
			color: var(--text-1);
		}
	}

	main {
		max-width: var(--size-lg);
		margin: 0 auto;
		padding: 0 var(--size-3);
	}
</style>
