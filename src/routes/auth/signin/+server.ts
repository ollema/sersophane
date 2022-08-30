import { json, type RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals }) => {
	const { email, password } = await request.json();

	const { token, user } = await locals.pocketbase.users.authViaEmail(email, password);

	return json({ token: token, user: user });
};
