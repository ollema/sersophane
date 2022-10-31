import type { PageServerLoad } from './$types';
import type { Event } from '$lib/types';

function parseFilterParams(urlParams: URLSearchParams) {
	const filters: string[] = [];

	const filterNameParam = urlParams.get('name');
	if (filterNameParam !== null) {
		filters.push(`name ~ "${filterNameParam}"`);
	}

	const venueFilterParams = urlParams.get('venue');
	if (venueFilterParams !== null) {
		filters.push(`venue ~ "${venueFilterParams}"`);
	}

	return filters.join(' && ');
}

export const load: PageServerLoad = async ({ locals, url }) => {
	const pageParam = parseInt(url.searchParams.get('page') || '');
	const page = isNaN(pageParam) ? undefined : pageParam;

	const perPageParam = parseInt(url.searchParams.get('per_page') || '');
	const perPage = isNaN(perPageParam) ? 15 : perPageParam;

	const filters = parseFilterParams(url.searchParams);
	const sort = url.searchParams.get('sort') || 'starts';

	const queryParams = {
		sort: sort,
		filter: filters,
		expand: 'venue.city, artists, responses(event).user'
	};

	const events = await locals.pb.collection('events').getList<Event>(page, perPage, queryParams);

	return {
		events: structuredClone(events),
		sort: sort
	};
};
