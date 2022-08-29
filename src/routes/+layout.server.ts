import type { Profile } from '$lib/types';
import type { LayoutServerLoad } from '../../.svelte-kit/types/src/routes/$types';
import { DateTime } from 'luxon';

export const load: LayoutServerLoad = async ({ locals }) => {
	console.log(DateTime.now().toFormat('HH:mm:ss'), 'load in +layout.server.ts triggered');

	let profile: Profile | undefined = undefined;

	if (locals.user && locals.user.profile) {
		profile = {
			id: locals.user.profile.id,
			name: locals.user.profile.name,
			avatar: locals.user.profile.avatar,
			created: locals.user.profile.created,
			updated: locals.user.profile.updated
		};
	}

	return {
		profile: profile
	};
};
