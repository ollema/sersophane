import type { RequestHandler } from '@sveltejs/kit';

import PocketBase from 'pocketbase';

import type { Venue } from '$lib/types';

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
			perPage = undefined;
		}
	} else {
		perPage = undefined;
	}

	const sortParam = urlParams.get('sort');
	const filterNameParam = urlParams.get('name');

	const queryParams: Record<string, string | number> = { expand: 'city' };
	if (sortParam !== null) {
		queryParams.sort = sortParam;
	}
	if (filterNameParam !== null) {
		queryParams.filter = `name ~ "${filterNameParam}"`;
	}

	const result = await client.records.getList('venues', page, perPage, queryParams);

	const venues: Venue[] = result.items.map((record) => {
		return {
			id: record.id,
			name: record.name,
			city: {
				id: record['@expand'].city.id,
				name: record['@expand'].city.name,
				slug: record['@expand'].city.slug
			},
			url: record.url,
			slug: record.slug
		};
	});

	return {
		body: {
			venues: venues
		}
	};
};
