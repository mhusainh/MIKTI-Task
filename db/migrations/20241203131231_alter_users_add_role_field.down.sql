BEGIN;

ALTER TABLE users
DROP COLUMN role;

COMMIT;