import type { RequestHandler } from '@sveltejs/kit';

import { getEvents } from '$lib/api';

export const GET: RequestHandler = async ({ url }) => {
	const { events, eventResponseMap, page, perPage, totalItems, sort } = await getEvents(url);

	return {
		body: {
			events: events,
			eventResponseMap: eventResponseMap,
			currentPage: page,
			perPage: perPage,
			totalItems: totalItems,
			sort: sort
		}
	};
};
