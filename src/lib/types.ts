export type City = {
	id: string;
	name: string;
	slug: string;
};

export type Venue = {
	id: string;
	name: string;
	city: City;
	url?: string;
	slug: string;
};
