import { json, type RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals }) => {
	const { email, password } = await request.json();

	const { token, user } = await locals.pocketbase.users.authViaEmail(email, password);

	console.log('signing in - calling `await invalidateAll()` - should trigger `load`');
	return json({ token: token, user: user });
};
