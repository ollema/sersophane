<script lang="ts">
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<svelte:head>
	<title>sersophane • sign in</title>
	<meta name="description" content="sign in to sersophane" />
</svelte:head>

<div>
	<ul>
		{#each data.methods || [] as { authUrl, name }}
			<li>
				<a href={authUrl}>{name}</a>
			</li>
		{/each}
	</ul>
</div>

<style lang="postcss">
	div {
		display: flex;
		justify-content: center;
		align-items: center;

		padding: var(--size-8);
	}
	ul {
		list-style: none;
	}
</style>
