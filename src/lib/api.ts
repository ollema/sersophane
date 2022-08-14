import PocketBase, { Record as PocketBaseRecord } from 'pocketbase';

import type { Event, EventResponse, Venue } from '$lib/types';

const client = new PocketBase('http://127.0.0.1:8090');

const defaultEventQueryParams: Record<string, string | null> = { expand: 'venue.city, artists, responses.profile' };
const defaultVenueQueryParams: Record<string, string | null> = { expand: 'city' };

export async function getEvents(url: URL) {
	const { page, perPage, queryParams } = parseQueryParams(new URLSearchParams(url.searchParams));
	queryParams.expand = defaultEventQueryParams.expand;

	const result = await client.records.getList('events', page, perPage, queryParams);
	const events: Event[] = result.items.map(processEventRecord);
	const eventResponseMap = await getEventResponseMap(events);

	return {
		events: events,
		eventResponseMap: eventResponseMap,
		page: page,
		perPage: perPage,
		totalItems: result.totalItems,
		sort: queryParams.sort
	};
}

export async function getEvent(id: string) {
	const result = await client.records.getOne('events', id, defaultEventQueryParams);

	const event: Event = processEventRecord(result);
	const eventResponseMap = await getEventResponseMap([event]);
	const eventResponses = eventResponseMap[event.id] || [];

	return { event, eventResponses };
}

function parseQueryParams(urlParams: URLSearchParams) {
	let page: number | undefined;
	let perPage: number | undefined;
	const queryParams: Record<string, string | null> = {};

	const pageParam = urlParams.get('page');
	if (pageParam !== null) {
		page = parseInt(pageParam);
		if (isNaN(page)) {
			page = undefined;
		}
	} else {
		page = undefined;
	}

	const perPageParam = urlParams.get('perPage');
	if (perPageParam !== null) {
		perPage = parseInt(perPageParam);
		if (isNaN(perPage)) {
			perPage = 20;
		}
	} else {
		perPage = 20;
	}

	const sortParam = urlParams.get('sort');
	if (sortParam !== null) {
		queryParams.sort = sortParam;
	} else {
		queryParams.sort = null;
	}

	const filters: string[] = [];

	const filterNameParam = urlParams.get('name');
	if (filterNameParam !== null) {
		filters.push(`name ~ "${filterNameParam}"`);
	}

	const venueFilterParams = urlParams.getAll('venue');
	const venueFilters: string[] = [];
	for (const venueFilterParam of venueFilterParams) {
		venueFilters.push(`venue.name ~ "${venueFilterParam}"`);
	}
	if (venueFilters.length > 0) {
		filters.push(venueFilters.join(' || '));
	}

	queryParams.filter = filters.join(' && ');

	console.log(queryParams.filter);

	return { page: page, perPage: perPage, queryParams: queryParams };
}

function processEventRecord(record: PocketBaseRecord) {
	return {
		id: record.id,
		name: record.name,
		venue: {
			id: record['@expand'].venue.id,
			name: record['@expand'].venue.name,
			city: {
				id: record['@expand'].venue['@expand'].city.id,
				name: record['@expand'].venue['@expand'].city.name,
				created: record['@expand'].venue['@expand'].city.created,
				updated: record['@expand'].venue['@expand'].city.updated
			},
			url: record['@expand'].venue.url,
			created: record['@expand'].venue.created,
			updated: record['@expand'].venue.updated
		},
		artists: record['@expand'].artists.map((artist: PocketBaseRecord) => {
			return {
				id: artist.id,
				name: artist.name,
				url: artist.url,
				created: artist.created,
				updated: artist.updated
			};
		}),
		type: record.type,
		cancelled: record.cancelled,
		starts: record.starts,
		ends: record.ends,
		url: record.url,
		created: record.created,
		updated: record.updated
	};
}

async function getEventResponseMap(events: Event[]) {
	const eventResponseMap: { [eventId: string]: EventResponse[] } = {};

	const filter = events
		.map((event) => {
			return `event.id = "${event.id}"`;
		})
		.join(' || ');

	const result = await client.records.getFullList('event_responses', undefined, { filter: filter, expand: 'event, profile' });
	const eventResponses: EventResponse[] = result.map(processResponseRecord);

	for (const eventResponse of eventResponses) {
		if (!eventResponseMap[eventResponse.event.id]) {
			eventResponseMap[eventResponse.event.id] = [];
		}
		eventResponseMap[eventResponse.event.id].push(eventResponse);
	}

	return eventResponseMap;
}

function processResponseRecord(record: PocketBaseRecord) {
	return {
		id: record.id,
		response: record.response,
		event: {
			id: record['@expand'].event.id
		},
		profile: {
			id: record['@expand'].profile.id,
			name: record['@expand'].profile.name,
			// avatar: `http://127.0.0.1:8090/api/files/profiles/${record.id}/${record['@expand'].profile.avatar}`,
			avatar: client.records.getFileUrl(record['@expand'].profile, record['@expand'].profile.avatar, { thumb: '100x100' }),
			created: record['@expand'].profile.created,
			updated: record['@expand'].profile.updated
		},
		created: record.created,
		updated: record.updated
	};
}

export async function getAllVenues() {
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
