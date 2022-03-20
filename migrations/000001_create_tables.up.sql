BEGIN;

CREATE COLLATION IF NOT EXISTS swedish (locale='sv_SE.utf8');

CREATE TYPE event_type AS enum (
    'concert',
    'festival',
    'film'
);

CREATE TABLE IF NOT EXISTS events (
    event_id bigserial PRIMARY KEY,
    event_name text NOT NULL COLLATE swedish,
    event_type event_type NOT NULL,
    event_created_at timestamp(0) WITH time zone NOT NULL DEFAULT now(),
    event_start timestamp(0) WITH time zone NOT NULL,
    event_end timestamp(0) WITH time zone NOT NULL,
    event_cancelled bool NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS artists (
    artist_id bigserial PRIMARY KEY,
    artist_name citext UNIQUE NOT NULL COLLATE swedish
);

CREATE TABLE IF NOT EXISTS venues (
    venue_id bigserial PRIMARY KEY,
    venue_name citext UNIQUE NOT NULL COLLATE swedish,
    venue_city citext NOT NULL COLLATE swedish
);

CREATE TABLE IF NOT EXISTS users (
    user_id bigserial PRIMARY KEY,
    user_name citext UNIQUE NOT NULL COLLATE swedish,
    user_created_at timestamp(0) WITH time zone NOT NULL DEFAULT now(),
    user_email citext UNIQUE NOT NULL,
    user_password_hash bytea NOT NULL,
    user_activated bool NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS event_artist (
    event_id bigint NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    artist_id bigint NOT NULL REFERENCES artists (artist_id) ON DELETE CASCADE,
    event_artist_running_order integer NOT NULL,
    CONSTRAINT event_artist_key PRIMARY KEY (event_id, artist_id)
);

CREATE TABLE IF NOT EXISTS event_venue (
    event_id bigint NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    venue_id bigint NOT NULL REFERENCES venues (venue_id) ON DELETE CASCADE,
    CONSTRAINT event_venue_key PRIMARY KEY (event_id, venue_id)
);

CREATE TYPE event_user_status AS enum (
    'interested',
    'going'
);

CREATE TABLE IF NOT EXISTS event_user (
    event_id bigint NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    user_id bigint NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    event_user_status event_user_status NOT NULL,
    CONSTRAINT event_user_key PRIMARY KEY (event_id, user_id)
);

-- Add common venues
INSERT INTO venues (venue_name, venue_city)
VALUES
    ('Bengans', 'Göteborg'),
    ('Bio Roy', 'Göteborg'),
    ('Brewhouse', 'Göteborg'),
    ('DDRC', 'Göteborg'),
    ('Fängelset', 'Göteborg'),
    ('Liseberg', 'Göteborg'),
    ('Musikens Hus', 'Göteborg'),
    ('M/S Götapetter', 'Göteborg'),
    ('Nefertiti', 'Göteborg'),
    ('Oceanen', 'Göteborg'),
    ('Pustervik', 'Göteborg'),
    ('Skjul Fyra Sex', 'Göteborg'),
    ('Sticky Fingers', 'Göteborg'),
    ('Studio HKPSM', 'Göteborg'),
    ('The Abyss', 'Göteborg'),
    ('Trädgårn', 'Göteborg'),
    ('Truckstop Alaska', 'Göteborg'),
    ('Ullevi', 'Göteborg'),
    ('Valand', 'Göteborg'),

    ('Annexet', 'Stockholm'),
    ('Avicii Arena', 'Stockholm'),
    ('Berns', 'Stockholm'),
    ('Circus', 'Stockholm'),
    ('Debaser', 'Stockholm'),
    ('Fryshuset', 'Stockholm'),
    ('Fållan', 'Stockholm'),
    ('Kraken', 'Stockholm'),
    ('Nalen', 'Stockholm'),
    ('Slaktkyrkan', 'Stockholm'),

    ('Plan B', 'Malmö');

-- Add common artists
INSERT INTO artists (artist_name)
VALUES
    ('Bombus'),
    ('Graveyard'),
    ('Gösta Berlings Saga'),
    ('Hammers of Misfortune'),
    ('Hällas'),
    ('OM'),
    ('Skraeckoedlan'),
    ('Sleep'),
    ('Vampire'),
    ('Vastum'),
    ('YOB');

COMMIT;
