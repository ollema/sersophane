BEGIN;

CREATE COLLATION swedish (locale='sv_SE.utf8');

CREATE TYPE event_type AS enum (
    'concert',
    'festival',
    'film'
);

CREATE TABLE IF NOT EXISTS events (
    id bigserial PRIMARY KEY,
    name text NOT NULL COLLATE swedish,
    type event_type NOT NULL,
    created_at timestamp(0) WITH time zone NOT NULL DEFAULT now(),
    start_at timestamp(0) WITH time zone NOT NULL,
    end_at timestamp(0) WITH time zone NOT NULL,
    cancelled bool NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS artists (
    id bigserial PRIMARY KEY,
    name citext UNIQUE NOT NULL COLLATE swedish
);

CREATE TABLE IF NOT EXISTS venues (
    id bigserial PRIMARY KEY,
    name citext UNIQUE NOT NULL COLLATE swedish,
    city citext NOT NULL COLLATE swedish
);

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name citext UNIQUE NOT NULL COLLATE swedish,
    created_at timestamp(0) WITH time zone NOT NULL DEFAULT now(),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS event_artist (
    event_id bigint NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    artist_id bigint NOT NULL REFERENCES artists (id) ON DELETE CASCADE,
    CONSTRAINT event_artist_key PRIMARY KEY (event_id, artist_id)
);

CREATE TABLE IF NOT EXISTS event_venue (
    event_id bigint NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    venue_id bigint NOT NULL REFERENCES venues (id) ON DELETE CASCADE,
    CONSTRAINT event_venue_key PRIMARY KEY (event_id, venue_id)
);

CREATE TYPE event_user_status AS enum (
    'interested',
    'going'
);

CREATE TABLE IF NOT EXISTS event_user (
    event_id bigint NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    user_id bigint NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    event_user_status event_user_status NOT NULL,
    CONSTRAINT event_user_key PRIMARY KEY (event_id, user_id)
);

-- Add common venues
INSERT INTO venues (name, city)
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
    ('Truckstop Alaska (RIP)', 'Göteborg'),
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
INSERT INTO artists (name)
VALUES
    ('Bombus'),
    ('Graveyard'),
    ('Gösta Berlings Saga'),
    ('Hällas'),
    ('Skraeckoedlan'),
    ('Vampire'),
    ('YOB');

COMMIT;
