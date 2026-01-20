BEGIN;

ALTER TABLE ctfs
    ADD COLUMN scoreboard_freeze_start TIMESTAMPTZ;
    ADD COLUMN scoreboard_freeze_end TIMESTAMPTZ;

END;
