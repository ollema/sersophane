import PocketBase, { Record as PocketBaseRecord } from 'pocketbase';

import type { Event } from '$lib/types';

const client = new PocketBase('http://127.0.0.1:8090');

const defaultQueryParams: Record<string, string | null> = { expand: 'venue.city, artists, responses.profile' };

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

	const filterNameParam = urlParams.get('name');
	if (filterNameParam !== null) {
		queryParams.filter = `name ~ "${filterNameParam}"`;
	}

	return { page: page, perPage: perPage, queryParams: queryParams };
}

function parseEventRecord(record: PocketBaseRecord) {
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
		responses: (record['@expand'].responses || []).map((resp: PocketBaseRecord) => {
			return {
				id: resp.id,
				response: resp.response,
				profile: {
					id: resp['@expand'].profile.id,
					name: resp['@expand'].profile.name,
					avatar: resp['@expand'].profile.avatar,
					created: resp['@expand'].profile.created,
					updated: resp['@expand'].profile.updated
				},
				created: resp.created,
				updated: resp.updated
			};
		}),
		url: record.url,
		created: record.created,
		updated: record.updated
	};
}

export async function getEvents(url: URL) {
	const { page, perPage, queryParams } = parseQueryParams(new URLSearchParams(url.searchParams));
	queryParams.expand = defaultQueryParams.expand;

	const result = await client.records.getList('events', page, perPage, queryParams);
	const events: Event[] = result.items.map(parseEventRecord);

	return { events: events, page: page, perPage: perPage, totalItems: result.totalItems, sort: queryParams.sort };
}

export async function getEvent(id: string) {
	const result = await client.records.getOne('events', id, defaultQueryParams);

	const event: Event = parseEventRecord(result);

	return event;
}
