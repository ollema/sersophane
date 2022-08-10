import type { RequestHandler } from '@sveltejs/kit';

import PocketBase from 'pocketbase';

import type { Event } from '$lib/types';

const client = new PocketBase('http://127.0.0.1:8090');

export const GET: RequestHandler = async ({ url }) => {
	const urlParams = new URLSearchParams(url.searchParams);

	const pageParam = urlParams.get('page');
	let page: number | undefined;
	if (pageParam !== null) {
		page = parseInt(pageParam);
		if (isNaN(page)) {
			page = undefined;
		}
	} else {
		page = undefined;
	}

	const perPageParam = urlParams.get('perPage');
	let perPage: number | undefined;
	if (perPageParam !== null) {
		perPage = parseInt(perPageParam);
		if (isNaN(perPage)) {
			perPage = 20;
		}
	} else {
		perPage = 20;
	}

	const sortParam = urlParams.get('sort');
	const filterNameParam = urlParams.get('name');

	const queryParams: Record<string, string | null> = { expand: 'venue, venue.city, artists' };
	if (sortParam !== null) {
		queryParams.sort = sortParam;
	} else {
		queryParams.sort = null;
	}
	if (filterNameParam !== null) {
		queryParams.filter = `name ~ "${filterNameParam}"`;
	}

	const result = await client.records.getList('events', page, perPage, queryParams);

	const events: Event[] = result.items.map((record) => {
		return {
			id: record.id,
			name: record.name,
			venue: {
				id: record['@expand'].venue.id,
				name: record['@expand'].venue.name,
				city: {
					id: record['@expand'].venue.city.id,
					name: record['@expand'].venue.city.name,
					slug: record['@expand'].venue.city.slug
				},
				url: record['@expand'].venue.url,
				slug: record['@expand'].venue.slug
			},
			artists: record['@expand'].artists.map((artist: Record<string, string | null>) => {
				return {
					id: artist.id,
					name: artist.name,
					url: artist.url,
					slug: artist.slug
				};
			}),
			startsAt: record.startsAt,
			endsAt: record.endsAt,
			type: record.type,
			url: record.url,
			slug: record.slug
		};
	});

	return {
		body: {
			events: events,
			currentPage: result.page,
			perPage: result.perPage,
			totalItems: result.totalItems,
			sort: queryParams.sort
		}
	};
};
