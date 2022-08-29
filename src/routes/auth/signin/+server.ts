import { json, type RequestHandler } from '@sveltejs/kit';
import { DateTime } from 'luxon';

export const POST: RequestHandler = async ({ request, locals }) => {
	console.log(DateTime.now().toFormat('HH:mm:ss'), 'signing in');

	const { email, password } = await request.json();

	const { token, user } = await locals.pocketbase.users.authViaEmail(email, password);

	return json({ token: token, user: user });
};
