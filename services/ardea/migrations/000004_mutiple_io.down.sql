DO language plpgsql $$
BEGIN;
    RAISE EXCEPTION 'Migration back to single dimension is not possible';
END;
$$;
