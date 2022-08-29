import { json, type RequestHandler } from '@sveltejs/kit';
import { DateTime } from 'luxon';

export const POST: RequestHandler = async ({ locals }) => {
	console.log(DateTime.now().toFormat('HH:mm:ss'), 'signing out');

	locals.pocketbase.authStore.clear();

	return json({});
};
