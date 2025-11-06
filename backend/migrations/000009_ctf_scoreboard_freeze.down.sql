BEGIN;

ALTER TABLE ctfs
    DROP COLUMN scoreboard_freeze_end;
    DROP COLUMN scoreboard_freeze_start;

END;
