BEGIN;

ALTER TABLE containers ADD auth jsonb DEFAULT '{"auths": {}}'::jsonb;

COMMIT;