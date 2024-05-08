BEGIN;

ALTER TABLE challtools_import_token ADD COLUMN ctf_event_id UUID REFERENCES ctf_events(id) ON DELETE CASCADE;
ALTER TABLE challtools_import_token DROP COLUMN name;
ALTER TABLE challtools_import_token DROP COLUMN last_used;

END;
