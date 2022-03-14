BEGIN;

CREATE TABLE IF NOT EXISTS models (
    id uuid primary key,
    owner_id uuid not null,
    state text not null,
    input_shape integer[] not null,
    output_shape integer[] not null,
    path text
);

CREATE INDEX IF NOT EXISTS  models_owner_id ON models USING btree (owner_id);

COMMIT;
