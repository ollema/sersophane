import { error, redirect } from '@sveltejs/kit';

import { AUTH_PROVIDER, AUTH_REDIRECT_URL } from '$env/static/private';

import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ locals, url }) => {
	const code = url.searchParams.get('code');

	if (code) {
		const { authProviders } = await locals.pb.collection('users').listAuthMethods();
		// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
		const { name, codeVerifier } = authProviders.find((provider) => provider.name === AUTH_PROVIDER)!;

		const { record, meta } = await locals.pb.collection('users').authWithOAuth2(name, code, codeVerifier, AUTH_REDIRECT_URL);

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
