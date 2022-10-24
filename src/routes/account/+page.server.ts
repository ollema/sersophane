import { redirect } from '@sveltejs/kit';
import type { Action, Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ parent }) => {
	return await parent();
};

const signout: Action = async ({ cookies, locals }) => {
	locals.pb.authStore.clear();
	cookies.delete('pb_auth');

	throw redirect(302, '/');
};

export const actions: Actions = { signout };
