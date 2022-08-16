<script>
	import { post } from '$lib/utils';
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';

	async function submit() {
		const response = await post('auth/login', { email, password });
		console.log(response);
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
