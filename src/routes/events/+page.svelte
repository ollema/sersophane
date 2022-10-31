<script lang="ts">
	import Table from './Table.svelte';

	import Filters from './Filters.svelte';
	import Pagination from './Pagination.svelte';

	import type { PageData } from './$types';

	export let data: PageData;
	$: ({ events, sort } = data);
</script>

<svelte:head>
	<title>sersophane • events</title>
	<meta name="description" content="list of events" />
</svelte:head>

<Filters />
<Table {events} {sort} />
<Pagination {events} />
