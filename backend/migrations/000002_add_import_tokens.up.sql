BEGIN;

CREATE TABLE challtools_import_token (
    id UUID NOT NULL PRIMARY KEY,
    token TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    ctf_event_id UUID REFERENCES ctf_events(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

END;
