<script lang="ts">
	import { page } from '$app/stores';
	import Nav from '$lib/components/Nav.svelte';
	import '../app.css';
	import '../reset.css';
	import 'tippy.js/dist/tippy.css';

	$: section = $page.url.pathname.split('/')[1];
</script>

<header>
	<Nav {section} />
</header>

<slot />

<style>
	:global(html) {
		font-family: 'Inter', 'system-ui';
		color: var(--text);
		background-color: var(--bg);
	}

	:global(body) {
		height: 100vh;

		display: grid;
		grid-template-columns: 1fr;
		grid-template-rows: 4rem 1fr 3rem;
		grid-template-areas:
			'header'
			'main'
			'footer';
	}

	:global(header) {
		grid-area: 'header';
	}

	:global(main) {
		grid-area: 'main';
		overflow: auto;

		display: flex;
		flex-direction: column;

		width: 100%;
		max-width: 1080px;

		margin: 0 auto;
		padding: 0.5rem;
	}

	:global(footer) {
		grid-area: 'footer';

		border-top: 3px solid var(--bg-secondary);
	}
</style>
