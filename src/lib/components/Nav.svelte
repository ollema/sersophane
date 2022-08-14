<script lang="ts">
	export let section: string;
</script>

<nav>
	<div>
		<div class="left-nav">
			<a sveltekit:prefetch href="/" class:selected={section === ''}>sersophane</a>
			<a sveltekit:prefetch href="/events" class:selected={section === 'events'}>events</a>
			<a sveltekit:prefetch href="/about" class:selected={section === 'about'}>about</a>
		</div>

		<a sveltekit:prefetch href="/" class:selected={section === 'profile'}>user: ...</a>
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

	a {
		padding: 0.75em 0.5em;

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
