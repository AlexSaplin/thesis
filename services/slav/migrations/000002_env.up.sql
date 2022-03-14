BEGIN;

ALTER TABLE containers ADD env jsonb;

COMMIT;