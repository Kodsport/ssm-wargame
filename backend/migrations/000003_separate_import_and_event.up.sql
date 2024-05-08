BEGIN;

ALTER TABLE challtools_import_token DROP COLUMN ctf_event_id;
ALTER TABLE challtools_import_token ADD COLUMN name TEXT NOT NULL;
ALTER TABLE challtools_import_token ADD COLUMN last_used TIMESTAMPTZ NOT NULL;

END;
