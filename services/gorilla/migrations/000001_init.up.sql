BEGIN;

CREATE TABLE IF NOT EXISTS deltas (
    id serial primary key,
    owner_id uuid not null,
    model_id uuid not null,
    category text not null,
    date timestamp not null,
    balance double precision not null
);

CREATE UNIQUE INDEX IF NOT EXISTS deltas_date_owner ON deltas USING btree (date, owner_id, model_id, category);

COMMIT;