import type { Handle } from '@sveltejs/kit';
import PocketBase, { User } from 'pocketbase';
import cookie from 'cookie';

export const handle: Handle = async ({ event, resolve }) => {
	const client = new PocketBase('http://127.0.0.1:8090');
	event.locals.pocketbase = client;

	const { token, user } = cookie.parse(event.request.headers.get('Cookie') ?? '');

	if (!token || !user) {
		return await resolve(event);
	}

	client.authStore.save(token, JSON.parse(user));

	return await resolve(event);
};
