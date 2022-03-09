BEGIN;

CREATE TYPE event_type AS enum (
    'concert',
    'festival',
    'film'
);

CREATE TABLE IF NOT EXISTS events (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    event_type event_type NOT NULL,
    created_at timestamp(0) WITH time zone NOT NULL DEFAULT now(),
    start_at timestamp(0) WITH time zone NOT NULL,
    end_at timestamp(0) WITH time zone NOT NULL,
    venue text NOT NULL
);

CREATE TABLE IF NOT EXISTS artists (
    id bigserial PRIMARY KEY,
    name citext UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name citext UNIQUE NOT NULL,
    created_at timestamp(0) WITH time zone NOT NULL DEFAULT now(),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    active bool NOT NULL
);

CREATE TABLE IF NOT EXISTS event_artist (
    event_id bigint NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    artist_id bigint NOT NULL REFERENCES artists (id) ON DELETE CASCADE,
    CONSTRAINT event_artist_key PRIMARY KEY (event_id, artist_id)
);

CREATE TABLE IF NOT EXISTS event_user (
    event_id bigint NOT NULL REFERENCES events (id) ON DELETE CASCADE,
    user_id bigint NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT event_user_key PRIMARY KEY (event_id, user_id)
);

COMMIT;
