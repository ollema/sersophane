import type { PageServerLoad, Actions } from './$types';

import type { Event } from '$lib/types';

import { error } from '@sveltejs/kit';

import type PocketBase from 'pocketbase';
import { ClientResponseError } from 'pocketbase';

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

export const load = (async ({ locals, url }) => {
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
}) satisfies PageServerLoad;

async function respond(pb: PocketBase, request: Request, response: 'going' | 'interested' | 'not-going') {
	const data = await request.formData();
	const eventId = data.get('id');
	const userId = pb.authStore.model?.id;
	if (userId === undefined) {
		return;
	}

	let record;
	try {
		record = await pb.collection('responses').getFirstListItem(`event="${eventId}" && user="${userId}"`);
		record = await pb.collection('responses').update(record.id, {
			response: response
		});
	} catch (e: unknown) {
		if (e instanceof ClientResponseError) {
			if (e.data.message === "The requested resource wasn't found.")
				record = await pb.collection('responses').create({
					response: response,
					event: eventId,
					user: userId
				});
		} else {
			throw error(500, JSON.stringify(e));
		}
	}

	return {
		status: 200,
		body: structuredClone(record)
	};
}

export const actions: Actions = {
	createEvent: async ({ locals, request }) => {
		const data = await request.formData();
		const event = await locals.pb.collection('events').create({
			name: data.get('name'),
			description: data.get('description'),
			starts: data.get('starts'),
			ends: data.get('ends'),
			venue: data.get('venue'),
			artists: data.getAll('artists')
		});

		return {
			status: 201,
			body: structuredClone(event)
		};
	},
	respondGoing: async ({ locals, request }) => {
		await respond(locals.pb, request, 'going');
	},
	respondInterested: async ({ locals, request }) => {
		await respond(locals.pb, request, 'interested');
	},
	respondNotGoing: async ({ locals, request }) => {
		await respond(locals.pb, request, 'not-going');
	}
};
