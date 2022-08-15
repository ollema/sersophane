<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<main>
	<table>
		<tr>
			<td>id</td>
			<td>{data.event.id}</td>
		</tr>
		<tr>
			<td>name</td>
			<td>{data.event.name}</td>
		</tr>
		<tr>
			<td>venue - id</td>
			<td>{data.event.venue.id}</td>
		</tr>
		<tr>
			<td>venue - name</td>
			<td>{data.event.venue.name}</td>
		</tr>
		<tr>
			<td>venue - url</td>
			<td>{data.event.venue.url}</td>
		</tr>
		<tr>
			<td>venue - city - id</td>
			<td>{data.event.venue.city.id}</td>
		</tr>
		<tr>
			<td>venue - city - name</td>
			<td>{data.event.venue.city.name}</td>
		</tr>
		{#each data.event.artists as artist, i}
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
			<td>{data.event.type}</td>
		</tr>
		<tr>
			<td>cancelled</td>
			<td>{data.event.cancelled}</td>
		</tr>
		<tr>
			<td>starts</td>
			<td>{data.event.starts}</td>
		</tr>
		<tr>
			<td>ends</td>
			<td>{data.event.ends}</td>
		</tr>
		{#each data.eventResponses as response, i}
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
			<td>{data.event.url}</td>
		</tr>
	</table>
</main>
