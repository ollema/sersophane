<script lang="ts">
	import { page } from '$app/stores';
</script>

<svelte:head>
	<title>sersophane • error: {$page.status}</title>
</svelte:head>

<main>
	<h1>{$page.status}</h1>
	{#if $page.error}
		<h2>{$page.error.message}</h2>
	{/if}
</main>

<footer />

<style lang="postcss">
	h2 {
		@apply pt-2;
	}
</style>
