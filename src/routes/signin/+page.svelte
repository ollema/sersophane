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

<form on:submit|preventDefault={submit}>
	<fieldset>
		<input type="email" required placeholder="Email" autocomplete="email" bind:value={email} />
	</fieldset>
	<fieldset>
		<input type="password" required placeholder="Password" autocomplete="current-password" bind:value={password} />
	</fieldset>
	<button type="submit"> Sign in </button>
</form>
