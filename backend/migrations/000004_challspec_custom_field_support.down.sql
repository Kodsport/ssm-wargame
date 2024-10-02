BEGIN;

ALTER TABLE challenges DROP COLUMN custom;
ALTER TABLE challenges DROP COLUMN chall_namespace;

END;
