import type { PageServerLoad } from './$types';

import { getEvents } from './events';
import { getAllVenues } from './venues';

export const load: PageServerLoad = async ({ locals, url }) => {
	const { events, eventResponseMap, page, perPage, totalItems, sortBy } = await getEvents(locals.pocketbase, url);
	const allVenues = await getAllVenues(locals.pocketbase);

	return {
		events: events,
		eventResponseMap: eventResponseMap,
		allVenues: allVenues,
		currentPage: page,
		perPage: perPage,
		totalItems: totalItems,
		sortBy: sortBy
	};
};
