BEGIN;

DROP INDEX deltas_date_owner;

ALTER TABLE deltas
RENAME COLUMN model_id TO object_id;

ALTER TABLE deltas
ADD object_type text;

UPDATE deltas
SET object_type = 'MODEL'
WHERE balance <= 0;

UPDATE deltas
SET object_type = 'UNKNOWN'
WHERE balance > 0;

ALTER TABLE deltas
ALTER COLUMN object_type SET NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS deltas_date_owner ON deltas USING btree (date, owner_id, object_id, object_type, category);

COMMIT;
