<script>
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';

	async function submit() {
		await fetch('auth/signin', {
			method: 'POST',
			body: JSON.stringify({ email, password }),
			headers: {
				'Content-Type': 'application/json'
			}
		});
		goto('/');
	}
</script>

<main>
	<form on:submit|preventDefault={submit}>
		<input bind:value={email} id="email" type="email" required placeholder="email" autocomplete="email" />
		<input bind:value={password} id="password" type="password" required placeholder="password" autocomplete="current-password" />
		<button type="submit"> sign in </button>
		<div class="divider">- or -</div>
		<a class="register" href="/">register</a>
	</form>
</main>

<style>
	form {
		margin-top: 3rem;

		display: flex;
		align-items: center;

		flex-direction: column;
	}

	input {
		width: 100%;
		max-width: 20rem;

		margin-bottom: 1.5rem;

		border-bottom: 3px solid var(--bg-secondary);
	}

	input:last-of-type {
		margin-bottom: 2rem;
	}

	input:hover,
	input:focus {
		border-bottom: 3px solid var(--fg);
	}

	input::placeholder {
		color: var(--text-darker);
	}

	button,
	a {
		width: 100%;
		max-width: 20rem;

		padding: 0.5rem 0.75rem;

		background-color: var(--bg-secondary);
		text-align: center;
	}

	.divider {
		margin: 0.5rem 0;

		color: var(--text-darker);
	}
</style>
