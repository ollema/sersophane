import { json, type RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ locals }) => {
	locals.pocketbase.authStore.clear();

	console.log('signing out - calling `await invalidateAll()` - should trigger `load`');
	return json({});
};
