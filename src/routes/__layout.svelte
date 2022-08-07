<script lang="ts">
	import { page } from '$app/stores';
	import Nav from '$lib/components/Nav.svelte';
	import '../app.css';

	$: section = $page.url.pathname.split('/')[1];
</script>

<Nav {section} />

<main>
	<slot />
</main>

<style>
	main {
		display: flex;
		flex-direction: column;

		width: 100%;
		max-width: 1080px;
		margin: 0 auto;
	}
</style>
