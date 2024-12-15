BEGIN;

ALTER TABLE users
    ADD COLUMN reset_password_token VARCHAR(255) NULL,
    ADD COLUMN verify_email_token VARCHAR(255) NULL,
    ADD COLUMN is_verified INT(1) NOT NULL DEFAULT 0;

COMMIT;