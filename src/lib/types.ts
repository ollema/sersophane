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
	event: Partial<Event>;
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
	cancelled?: boolean;
	starts: string;
	ends: string;
	url?: string;

	created: string;
	updated: string;
};

// types for POST endpoints
export type CityPost = Omit<City, 'id' | 'created' | 'updated'>;

export type VenuePost = Omit<Venue, 'id' | 'created' | 'updated' | 'city'> & { cityId: string };

export type ArtistPost = Omit<Artist, 'id' | 'created' | 'updated'>;

export type ResponsePost = Omit<EventResponse, 'id' | 'created' | 'updated' | 'event' | 'profile'> & {
	eventId: string;
	profileId: string;
};

export type EventPost = Omit<Event, 'id' | 'created' | 'updated' | 'venue' | 'artists' | 'responses'> & {
	venueId: string;
	artistIds: string[];
};
