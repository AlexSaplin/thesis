BEGIN;

ALTER TABLE containers DROP COLUMN auth;

COMMIT;