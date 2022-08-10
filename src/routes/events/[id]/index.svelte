<script lang="ts">
	import type { Event } from '$lib/types';

	export let event: Event;
</script>

<main>
	<table>
		<tr>
			<td>id</td>
			<td>{event.id}</td>
		</tr>
		<tr>
			<td>name</td>
			<td>{event.name}</td>
		</tr>
		<tr>
			<td>venue - id</td>
			<td>{event.venue.id}</td>
		</tr>
		<tr>
			<td>venue - name</td>
			<td>{event.venue.name}</td>
		</tr>
		<tr>
			<td>venue - url</td>
			<td>{event.venue.url}</td>
		</tr>
		<tr>
			<td>venue - city - id</td>
			<td>{event.venue.city.id}</td>
		</tr>
		<tr>
			<td>venue - city - name</td>
			<td>{event.venue.city.name}</td>
		</tr>
		{#each event.artists as artist, i}
			<tr>
				<td>artist {i + 1} - id</td>
				<td>{artist.id}</td>
			</tr>
			<tr>
				<td>artist {i + 1} - name</td>
				<td>{artist.name}</td>
			</tr>
			<tr>
				<td>artist - url</td>
				<td>{artist.url}</td>
			</tr>
		{/each}
		<tr>
			<td>type</td>
			<td>{event.type}</td>
		</tr>
		<tr>
			<td>cancelled</td>
			<td>{event.cancelled}</td>
		</tr>
		<tr>
			<td>startsAt</td>
			<td>{event.startsAt}</td>
		</tr>
		<tr>
			<td>endsAt</td>
			<td>{event.endsAt}</td>
		</tr>
		{#each event.responses as response, i}
			<tr>
				<td>response {i + 1} - id</td>
				<td>{response.id}</td>
			</tr>
			<tr>
				<td>response {i + 1} - response</td>
				<td>{response.response}</td>
			</tr>
			<tr>
				<td>response {i + 1} - profile - id</td>
				<td>{response.profile.id}</td>
			</tr>
			<tr>
				<td>response {i + 1} - profile - name</td>
				<td>{response.profile.name}</td>
			</tr>
			<tr>
				<td>response {i + 1} - profile - avatar</td>
				<td>{response.profile.avatar}</td>
			</tr>
		{/each}
		<tr>
			<td>url</td>
			<td>{event.url}</td>
		</tr>
	</table>
</main>
