<script lang="ts">
	import Table from './Table.svelte';

	import type { PageData } from './$types';
	import Filters from './Filters.svelte';

	export let data: PageData;
	$: ({ events, sort } = data);
</script>

<svelte:head>
	<title>sersophane • events</title>
	<meta name="description" content="overview of all upcoming and past events" />
</svelte:head>

<Filters />
<Table {events} {sort} />
