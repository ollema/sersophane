import PocketBase, { Record } from 'pocketbase';

import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.pb = new PocketBase('http://127.0.0.1:8090');
	event.locals.pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

	if (event.locals.pb.authStore.isValid) {
		const record = event.locals.pb.authStore.model as Record;
		event.locals.user = {
			id: record.id,
			username: record.username,
			name: record.name,
			avatar: record.avatar
		};
	}

	const response = await resolve(event);
	return response;
};
