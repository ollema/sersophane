import { json, type RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ locals }) => {
	locals.pocketbase.authStore.clear();

	return json({});
};
