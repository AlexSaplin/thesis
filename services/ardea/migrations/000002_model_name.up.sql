BEGIN;

ALTER TABLE models ADD COLUMN name text;

CREATE UNIQUE INDEX IF NOT EXISTS  models_owner_id_name ON models USING btree (owner_id, name);

COMMIT;
