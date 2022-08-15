import type { PageServerLoad } from './$types';

import { getAllVenues, getEvents } from '$lib/api';

export const load: PageServerLoad = async ({ url }) => {
	const { events, eventResponseMap, page, perPage, totalItems, sort } = await getEvents(url);
	const allVenues = await getAllVenues();

	return {
		events: events,
		eventResponseMap: eventResponseMap,
		allVenues: allVenues,
		currentPage: page,
		perPage: perPage,
		totalItems: totalItems,
		sort: sort
	};
};
