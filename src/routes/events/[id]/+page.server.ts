import type { PageServerLoad } from './$types';

import { getEvent } from '$lib/api';

export const load: PageServerLoad = async ({ params }) => {
	const { event, eventResponses } = await getEvent(params.id);

	return {
		event: event,
		eventResponses: eventResponses
	};
};
