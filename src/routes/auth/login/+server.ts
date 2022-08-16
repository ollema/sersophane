import type { RequestHandler } from '@sveltejs/kit';
import cookie from 'cookie';

const defaultCookieOptions = { maxAge: 30 * 24 * 60 * 60, path: '/', httpOnly: true, sameSite: true, secure: true };

export const POST: RequestHandler = async ({ request, locals: { pocketbase } }) => {
	const response = new Response('');

	const { email, password } = await request.json();
	try {
		const { token, user } = await pocketbase.users.authViaEmail(email, password);
		pocketbase.authStore.save(token, user);
		response.headers.append('Set-Cookie', cookie.serialize('token', token, defaultCookieOptions));
		response.headers.append('Set-Cookie', cookie.serialize('user', JSON.stringify(user), defaultCookieOptions));
	} catch {}

	return response;
};
