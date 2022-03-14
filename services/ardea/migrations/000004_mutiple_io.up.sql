BEGIN;

ALTER TABLE models ADD COLUMN input_shape_md jsonb DEFAULT '[]'::jsonb;
ALTER TABLE models ADD COLUMN output_shape_md jsonb DEFAULT '[]'::jsonb;

UPDATE models SET input_shape_md = '[[]]'::jsonb;
UPDATE models SET output_shape_md = '[[]]'::jsonb;

UPDATE models SET input_shape_md = jsonb_set(input_shape_md::jsonb, '{0}', array_to_json(input_shape)::jsonb);
UPDATE models SET output_shape_md = jsonb_set(output_shape_md::jsonb, '{0}', array_to_json(output_shape)::jsonb);

ALTER TABLE models DROP COLUMN input_shape;
ALTER TABLE models DROP COLUMN output_shape;

ALTER TABLE models RENAME COLUMN input_shape_md TO input_shape;
ALTER TABLE models RENAME COLUMN output_shape_md TO output_shape;

COMMIT;
