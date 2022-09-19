import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
	console.log('load triggered');
	if (locals.user && locals.user.profile) {
		return {
			profile: {
				id: locals.user.profile.id,
				name: locals.user.profile.name,
				avatar: locals.user.profile.avatar,
				created: locals.user.profile.created,
				updated: locals.user.profile.updated
			}
		};
	} else {
		return {
			profile: undefined
		};
	}
};
