import { AUTH_PROVIDER, AUTH_REDIRECT_URL } from '$env/static/private';
import { error, redirect } from '@sveltejs/kit';

import type { RequestHandler } from './$types';

export const GET = (async ({ locals, cookies }) => {
	const { authProviders } = await locals.pb.collection('users').listAuthMethods();
	const provider = authProviders.find((method) => method.name === AUTH_PROVIDER);
	if (!provider) {
		throw error(500, 'no auth provider found');
	}

	cookies.set('provider', JSON.stringify(provider), {
		path: '/',
		maxAge: 60 * 5 // 5 minutes
	});

	throw redirect(302, provider.authUrl + AUTH_REDIRECT_URL);
}) satisfies RequestHandler;
