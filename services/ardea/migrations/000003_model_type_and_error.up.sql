BEGIN;

ALTER TABLE models ADD COLUMN value_type text NOT NULL;
ALTER TABLE models ADD COLUMN err text;

COMMIT;
