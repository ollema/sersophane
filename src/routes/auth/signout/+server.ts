import type { RequestHandler } from '@sveltejs/kit';
import cookie from 'cookie';

const defaultCookieOptions = { maxAge: -1, path: '/', httpOnly: true, sameSite: true, secure: true };

export const POST: RequestHandler = async ({ locals }) => {
	const response = new Response('{}');

	locals.pocketbase.authStore.clear();
	locals.user = undefined;
	response.headers.append('Set-Cookie', cookie.serialize('token', '', defaultCookieOptions));
	response.headers.append('Set-Cookie', cookie.serialize('user', '', defaultCookieOptions));

	return response;
};
