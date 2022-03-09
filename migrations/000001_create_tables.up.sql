begin;

create type event_type as enum ('concert', 'festival', 'film');

create table if not exists events (
    id bigserial primary key,
    name text not null,
    type event_type not null,
    created timestamp(0) with time zone not null default now(),
    start timestamp(0) with time zone not null,
    end timestamp(0) with time zone not null,
    location text not null
);

create table if not exists artists (
    id bigserial primary key,
    name text not null
);

create table if not exists users (
    id bigserial primary key,
    name text not null
    created timestamp(0) with time zone not null default now(),
    active bool not null,
);

create table if not exists event_artist (
    event_id bigint not null references events (id) on delete cascade,
    artist_id bigint not null references artists (id) on delete cascade,
    constraint event_artist_key primary key (event_id, artist_id)
);

create table if not exists event_user (
    event_id bigint not null references events (id) on delete cascade,
    user_id bigint not null references users (id) on delete cascade,
    constraint event_user_key primary key (event_id, user_id)
);

commit;
