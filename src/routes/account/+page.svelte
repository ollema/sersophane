<script lang="ts">
	import type { PageServerData } from './$types';

	export let data: PageServerData;
	$: ({ user } = data);
</script>

<svelte:head>
	<title>sersophane • account</title>
	<meta name="description" content="account details for {user?.name}" />
</svelte:head>

<div class="account">
	<div class="details">
		<table>
			<tr>
				<td>username:</td>
				<td>{user?.username}</td>
			</tr>
			<tr>
				<td>display name:</td>
				<td>{user?.name}</td>
			</tr>
		</table>

		<form action="?/signout" method="POST">
			<button type="submit">sign out</button>
		</form>
	</div>
</div>

<style lang="postcss">
	.account {
		max-width: calc(var(--size-sm) + var(--size-12));
		margin: 0 auto;

		padding: var(--size-3);
		margin-top: var(--size-3);

		background-color: var(--surface-2);
		border: var(--border-size-2) solid var(--surface-3);
		border-radius: var(--radius-2);
	}

	table {
		display: grid;
		gap: var(--size-1) var(--size-2);

		grid-template-columns:
			auto
			minmax(2rem, 1fr);

		padding-bottom: var(--size-5);
	}

	tr {
		display: contents;
	}

	td {
		position: relative;
		display: flex;
		align-items: center;

		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	td:first-child {
		justify-content: flex-end;
	}

	form {
		display: flex;
		align-items: center;
		justify-content: center;
	}
</style>
