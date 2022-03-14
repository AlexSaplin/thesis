BEGIN;

CREATE TABLE IF NOT EXISTS containers (
    id uuid primary key,
    name text not null unique,
    scale integer not null,
    instance_type text not null,
    image text not null,
    owner_id text not null,
    port integer not null
);

COMMIT;