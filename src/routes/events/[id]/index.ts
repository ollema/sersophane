import type { RequestHandler } from '@sveltejs/kit';

import { getEvent } from '$lib/api';

export const GET: RequestHandler = async ({ params }) => {
	const { event, eventResponses } = await getEvent(params.id);

	return {
		body: {
			event: event,
			eventResponses: eventResponses
		}
	};
};
