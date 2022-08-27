import { getEvent } from '../events';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, params }) => {
	const { event, eventResponses } = await getEvent(locals.pocketbase, params.id);

	return {
		event: event,
		eventResponses: eventResponses
	};
};
