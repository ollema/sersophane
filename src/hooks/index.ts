import type { Handle } from '@sveltejs/kit';
import PocketBase from 'pocketbase';
import cookie from 'cookie';

export const handle: Handle = async ({ event, resolve }) => {
	const client = new PocketBase('http://127.0.0.1:8090');
	event.locals.pocketbase = client;

	const data = cookie.parse(event.request.headers.get('cookie') ?? '');
	if (!data.token || !data.model) {
		return await resolve(event);
	}

	try {
		const model = JSON.parse(data.model);
		client.authStore.save(data.token, model);
	} catch {
		client.authStore.clear();
	}

	return await resolve(event);
};
