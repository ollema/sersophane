import { error, redirect } from '@sveltejs/kit';

import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ locals, url }) => {
	const code = url.searchParams.get('code');

	if (code) {
		const { authProviders } = await locals.pb.collection('users').listAuthMethods();
		const { name, codeVerifier } = authProviders[0];

		const { record, meta } = await locals.pb
			.collection('users')
			.authWithOAuth2(name, code, codeVerifier, 'http://localhost:5173/auth/redirect');

		// get name/username from OAuth2 provider if not already set
		if (meta.name && record.name == '') {
			await locals.pb.collection('users').update(record.id, { name: meta.name });
		}
		if (meta.username && record.username == '') {
			await locals.pb.collection('users').update(record.id, { username: meta.username });
		}

		await locals.pb.collection('users').authRefresh();
		throw redirect(302, '/account');
	} else {
		throw error(404);
	}
};
