-- 000001_init.up.sql
CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version         BIGINT NOT NULL DEFAULT 1,
    full_name       VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    phone_number    VARCHAR(15)
);

CREATE TABLE todoapp.tasks (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title           VARCHAR(100) NOT NULL,
    completed       BOOLEAN NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL,
    completed_at    TIMESTAMPTZ,
    author_user_id  UUID NOT NULL REFERENCES todoapp.users(id)
);