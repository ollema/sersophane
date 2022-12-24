import type { PageServerLoad } from './$types';

export const load = (async ({ locals }) => {
	const allVenues: string[] = (await locals.pb.collection('venues').getFullList(100, { sort: 'name' })).map((venue) => venue.name);
	const allArtists: string[] = (await locals.pb.collection('artists').getFullList(100, { sort: 'name' })).map((artist) => artist.name);

	return {
		allVenues: allVenues,
		allArtists: allArtists
	};
}) satisfies PageServerLoad;
