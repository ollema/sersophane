import type { RequestHandler } from '@sveltejs/kit';

import { getAllVenues, getEvents } from '$lib/api';

export const GET: RequestHandler = async ({ url }) => {
	const { events, eventResponseMap, page, perPage, totalItems, sort } = await getEvents(url);
	const allVenues = await getAllVenues();

	return {
		body: {
			events: events,
			eventResponseMap: eventResponseMap,
			allVenues: allVenues,
			currentPage: page,
			perPage: perPage,
			totalItems: totalItems,
			sort: sort
		}
	};
};
