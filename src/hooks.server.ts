import type { Handle } from '@sveltejs/kit';
import PocketBase, { User } from 'pocketbase';

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.pocketbase = new PocketBase('http://127.0.0.1:8090');
	event.locals.pocketbase.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

	if (event.locals.pocketbase.authStore.isValid) {
		event.locals.user = event.locals.pocketbase.authStore.model as User;
	}

	const response = await resolve(event);

	// TODO: should probably enable secure cookies when deploying, ideally controlled with an env var
	response.headers.set('set-cookie', event.locals.pocketbase.authStore.exportToCookie({ secure: false }));

	return response;
};
