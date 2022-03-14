BEGIN;

DROP INDEX deltas_date_owner;

ALTER TABLE deltas
RENAME COLUMN object_id TO object_id;

ALTER TABLE deltas
DROP COLUMN object_type;

CREATE UNIQUE INDEX IF NOT EXISTS deltas_date_owner ON deltas USING btree (date, owner_id, model_id, category);

COMMIT;
