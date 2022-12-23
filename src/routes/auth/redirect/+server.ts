import { AUTH_REDIRECT_URL } from '$env/static/private';

import { error, redirect } from '@sveltejs/kit';

import type { RequestHandler } from './$types';
import { getTokenPayload } from 'pocketbase';

export const GET = (async ({ locals, url, cookies }) => {
	// parse the query parameters from the redirected url
	const params = url.searchParams;
	const state = params.get('state');
	if (!state) {
		throw error(404, 'no state query parameter found');
	}
	const code = params.get('code');
	if (!code) {
		throw error(404, 'no code query parameter found');
	}

	// load the previously stored provider's data from cookie
	const rawProvider = cookies.get('provider');
	if (!rawProvider) {
		throw error(404, 'no provider cookie found');
	}
	const provider = JSON.parse(rawProvider);

	// compare the redirect's state param and the stored provider's one
	if (state !== provider.state) {
		throw error(404, "state parameters don't match");
	}

	// authenticate the user with the provider
	const { record, meta } = await locals.pb
		.collection('users')
		.authWithOAuth2(provider.name, code, provider.codeVerifier, AUTH_REDIRECT_URL);

	// update the user's record with the data from the auth provider
	if (meta?.name && record.name == '') {
		await locals.pb.collection('users').update(record.id, { name: meta.name });
	}
	if (meta?.id) {
		await locals.pb.collection('users').update(record.id, { username: meta.id });
	}

	// refresh the auth store
	await locals.pb.collection('users').authRefresh();

	// clear the provider cookie
	cookies.delete('provider', { path: '/' });

	// set the pb_auth cookie
	const payload = getTokenPayload(locals.pb.authStore.token);
	cookies.set(
		'pb_auth',
		JSON.stringify({
			token: locals.pb.authStore.token,
			model: locals.pb.authStore.model
		}),
		{
			path: '/',
			expires: new Date(payload.exp * 1000)
		}
	);

	throw redirect(302, '/account');
}) satisfies RequestHandler;
