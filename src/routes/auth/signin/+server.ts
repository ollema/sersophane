import { AUTH_PROVIDER, AUTH_REDIRECT_URL } from '$env/static/private';
import { redirect } from '@sveltejs/kit';

import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ locals }) => {
	const { authProviders } = await locals.pb.collection('users').listAuthMethods();
	// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
	const { authUrl } = authProviders.find((provider) => provider.name === AUTH_PROVIDER)!;

	throw redirect(302, authUrl + AUTH_REDIRECT_URL);
};
