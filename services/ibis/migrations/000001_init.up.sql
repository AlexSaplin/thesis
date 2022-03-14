BEGIN;

CREATE TABLE IF NOT EXISTS functions (
    id uuid primary key,
    owner_id uuid not null,
    state text not null,
    code_path text,
    image text,
    name text,
    meta text,
    err text
);

CREATE INDEX IF NOT EXISTS  functions_owner_id ON functions USING btree (owner_id);

CREATE UNIQUE INDEX IF NOT EXISTS  functions_owner_id_name ON functions USING btree (owner_id, name);

COMMIT;
