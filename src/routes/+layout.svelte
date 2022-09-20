<script lang="ts">
	import { page } from '$app/stores';

	import Nav from './Nav.svelte';

	import '../app.css';
	import 'svooltip/svooltip.css';

	import type { LayoutServerData } from './$types';

	export let data: LayoutServerData;
	$: ({ profile } = data);

	$: section = $page.url.pathname.split('/')[1];
</script>

<header>
	<Nav {section} {profile} />
</header>

<slot />
