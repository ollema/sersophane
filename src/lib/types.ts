export type City = {
	id: string;
	name: string;

	created: string;
	updated: string;
};

export type Venue = {
	id: string;
	name: string;
	city: City;
	url?: string;

	created: string;
	updated: string;
};

export type Artist = {
	id: string;
	name: string;
	url?: string;

	created: string;
	updated: string;
};

export type Profile = {
	id: string;
	name: string;
	avatar: string;

	created: string;
	updated: string;
};

export enum Response {
	Interested = 'interested',
	Going = 'going',
	NotGoing = 'not-going'
}

export type EventResponse = {
	id: string;
	response: Response;
	profile: Profile;

	created: string;
	updated: string;
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
	type: EventType;
	cancelled: boolean;
	startsAt: string;
	endsAt: string;
	responses: EventResponse[];
	url?: string;

	created: string;
	updated: string;
};
