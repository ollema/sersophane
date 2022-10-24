import type { PageServerLoad } from './$types';
import type { Event } from '$lib/types';

function parsePageParams(urlParams: URLSearchParams) {
	let page: number | undefined;
	let perPage: number | undefined;

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

	return { page: page, perPage: perPage };
}

function parseFilterParams(urlParams: URLSearchParams) {
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

	return filters;
}

export const load: PageServerLoad = async ({ locals, url }) => {
	const { page, perPage } = parsePageParams(url.searchParams);
	const filters = parseFilterParams(url.searchParams);
	const sort = url.searchParams.get('sort');

	const queryParams: Record<string, string | null> = {
		sort: sort,
		filter: filters.join(' && '),
		expand: 'venue.city, artists, responses(event).user'
	};

	const events = await locals.pb.collection('events').getList<Event>(page, perPage, queryParams);

	return {
		events: structuredClone(events),
		sort: sort
	};
};
