export type IsoDateString = string;

export type RecordIdString = string;

export type UserIdString = string;

export type BaseRecord = {
	id: RecordIdString;
	created: IsoDateString;
	updated: IsoDateString;
};

export interface User extends BaseRecord {
	name: string;
	username: string;
	avatar: string;
}

export interface City extends BaseRecord {
	name: string;
}

export interface Venue extends BaseRecord {
	name: string;
	city: RecordIdString;
	url?: string;

	expand: {
		city: City;
	};
}

export interface Artist extends BaseRecord {
	name: string;
	url?: string;
}

export interface Response extends BaseRecord {
	response: 'interested' | 'going' | 'not-going';
	event: RecordIdString;
	user: RecordIdString;

	expand: {
		user: User;
	};
}

export interface Event extends BaseRecord {
	name: string;
	venue: RecordIdString;
	artists: RecordIdString[];
	type: 'concert' | 'music-festival' | 'movie';
	cancelled: boolean;
	starts: string;
	ends: string;
	url?: string;

	expand: {
		venue: Venue;
		artists: Artist[];
		'responses(event)': Response[];
	};
}
