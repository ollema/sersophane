import type PocketBase from 'pocketbase';
import type { Record as PocketBaseRecord } from 'pocketbase';

import type { Venue } from '$lib/types';

const defaultVenueQueryParams: Record<string, string | null> = { expand: 'city' };

export async function getAllVenues(client: PocketBase) {
	const result = await client.records.getFullList('venues', undefined, defaultVenueQueryParams);
	const venues: Venue[] = result.map(processVenueRecord);

	return venues;
}

function processVenueRecord(record: PocketBaseRecord) {
	return {
		id: record.id,
		name: record.name,
		city: {
			id: record['@expand'].city.id,
			name: record['@expand'].city.name,
			created: record['@expand'].city.created,
			updated: record['@expand'].city.updated
		},
		created: record.created,
		updated: record.updated
	};
}
