<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';

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
		await invalidateAll();
		goto('/');
	}
</script>

<svelte:head>
	<title>sersophane • sign in</title>
</svelte:head>

<main>
	<form on:submit|preventDefault={submit}>
		<input bind:value={email} id="email" type="email" required placeholder="email" autocomplete="email" />
		<input bind:value={password} id="password" type="password" required placeholder="password" autocomplete="current-password" />
		<button type="submit"> sign in </button>
		<div class="divider">- or -</div>
		<button type="submit">register</button>
	</form>
</main>

<footer />

<style lang="postcss">
	form {
		@apply flex flex-col items-center justify-center mt-12;
	}

	input {
		@apply w-full max-w-xs mt-2 bg-zinc-900 autofill:bg-zinc-900 border-0 border-b-2 border-zinc-800 focus:ring-0 focus:border-emerald-600 hover:border-emerald-300;
	}

	input:-webkit-autofill,
	input:-webkit-autofill:hover,
	input:-webkit-autofill:focus {
		-webkit-text-fill-color: #a1a1aa;
		caret-color: #a1a1aa;
		box-shadow: 0 0 0px 1000px #18181b inset;
		-webkit-box-shadow: 0 0 0px 1000px #18181b inset;
	}

	input::placeholder {
		@apply text-zinc-500;
	}

	button {
		@apply w-full max-w-xs py-2 px-3 bg-zinc-800 text-center first-of-type:mt-6;
	}

	.divider {
		@apply my-2 text-zinc-300;
	}
</style>
