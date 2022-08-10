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

export type Artist = {
	id: string;
	name: string;
	url?: string;
	slug: string;
};

export enum EventType {
	Concert = 'concert',
	Festival = 'festival',
	Film = 'film'
}

export type Event = {
	id: string;
	name: string;
	venue: Venue;
	artists: Artist[];
	startsAt: string;
	endsAt: string;
	type: EventType;
	url?: string;
	slug: string;
};
