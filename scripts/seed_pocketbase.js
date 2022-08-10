// silence experimental warnings
const originalEmit = process.emit;
process.emit = function (name, data) {
	if (name === `warning` && typeof data === `object` && data.name === `ExperimentalWarning`) return false;

	return originalEmit.apply(process, arguments);
};

// start of actual script
import PocketBase from 'pocketbase';

const pocketBaseUrl = 'http://127.0.0.1:8090';

console.log(`will seed database at ${pocketBaseUrl}`);

const client = new PocketBase(pocketBaseUrl);
const user = process.env.POCKET_BASE_ADMIN_USER;
const password = process.env.POCKET_BASE_ADMIN_PASS;

if (user === undefined || password === undefined) {
	throw Error('POCKET_BASE_ADMIN_USER or POCKET_BASE_ADMIN_PASS not set');
}

await client.admins.authViaEmail(user, password);

// create cities
const cities = [{ name: 'Göteborg' }, { name: 'Malmö' }, { name: 'Stockholm' }];

const createCities = async () => {
	for (let city of cities) {
		const record = await client.records.getList('cities', 1, 50, {
			filter: `name = "${city.name}"`
		});

		if (record.totalItems === 1) {
			await client.records.update('cities', record.items[0].id, city);
			continue;
		}

		if (record.totalItems === 0) {
			await client.records.create('cities', city);
			continue;
		}

		throw Error(`this should not happen - more than 1 city with the name ${city.name}`);
	}
};

await createCities();

// create venues
const venues = [
	{ name: 'Bengans', city: 'Göteborg' },
	{ name: 'Bio Roy', city: 'Göteborg' },
	{ name: 'Brewhouse', city: 'Göteborg' },
	{ name: 'DDRC', city: 'Göteborg' },
	{ name: 'Fängelset', city: 'Göteborg' },
	{ name: 'Liseberg', city: 'Göteborg' },
	{ name: 'Musikens Hus', city: 'Göteborg' },
	{ name: 'M/S Götapetter', city: 'Göteborg' },
	{ name: 'Nefertiti', city: 'Göteborg' },
	{ name: 'Oceanen', city: 'Göteborg' },
	{ name: 'Pustervik', city: 'Göteborg' },
	{ name: 'Skjul Fyra Sex', city: 'Göteborg' },
	{ name: 'Sticky Fingers', city: 'Göteborg' },
	{ name: 'Studio HKPSM', city: 'Göteborg' },
	{ name: 'The Abyss', city: 'Göteborg' },
	{ name: 'Trädgårn', city: 'Göteborg' },
	{ name: 'Truckstop Alaska', city: 'Göteborg' },
	{ name: 'Ullevi', city: 'Göteborg' },
	{ name: 'Valand', city: 'Göteborg' },

	{ name: 'Annexet', city: 'Stockholm' },
	{ name: 'Avicii Arena', city: 'Stockholm' },
	{ name: 'Berns', city: 'Stockholm' },
	{ name: 'Circus', city: 'Stockholm' },
	{ name: 'Debaser', city: 'Stockholm' },
	{ name: 'Fryshuset', city: 'Stockholm' },
	{ name: 'Fållan', city: 'Stockholm' },
	{ name: 'Kraken', city: 'Stockholm' },
	{ name: 'Nalen', city: 'Stockholm' },
	{ name: 'Slaktkyrkan', city: 'Stockholm' },

	{ name: 'Plan B', city: 'Malmö' }
];

const createVenues = async () => {
	for (let venue of venues) {
		const cityRecord = await client.records.getList('cities', 1, 50, {
			filter: `name = "${venue.city}"`
		});
		if (cityRecord.totalItems !== 1) {
			throw Error(`unknown city ${venue.city}`);
		}
		const cityId = cityRecord.items[0].id;
		venue.city = cityId;

		const record = await client.records.getList('venues', 1, 50, {
			filter: `name = "${venue.name}"`
		});

		if (record.totalItems === 1) {
			await client.records.update('venues', record.items[0].id, venue);
			continue;
		}

		if (record.totalItems === 0) {
			await client.records.create('venues', venue);
			continue;
		}

		throw Error(`this should not happen - more than 1 venue with the name ${venue.name}`);
	}
};

await createVenues();

// create artists
const artists = [
	{ name: 'Bombus' },
	{ name: 'Graveyard' },
	{ name: 'Gösta Berlings Saga' },
	{ name: 'Hammers of Misfortune' },
	{ name: 'Hällas' },
	{ name: 'OM' },
	{ name: 'Skraeckoedlan' },
	{ name: 'Sleep' },
	{ name: 'Vampire' },
	{ name: 'Vastum' },
	{ name: 'YOB' }
];

const createArtists = async () => {
	for (let artist of artists) {
		const record = await client.records.getList('artists', 1, 50, {
			filter: `name = "${artist.name}"`
		});

		if (record.totalItems === 1) {
			await client.records.update('artists', record.items[0].id, artist);
			continue;
		}

		if (record.totalItems === 0) {
			await client.records.create('artists', artist);
			continue;
		}

		throw Error(`this should not happen - more than 1 artist with the name ${artist.name}`);
	}
};

await createArtists();
