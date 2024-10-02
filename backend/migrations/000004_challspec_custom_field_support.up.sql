BEGIN;

ALTER TABLE challenges ADD COLUMN custom JSONB;
ALTER TABLE challenges ADD COLUMN chall_namespace TEXT;

END;
