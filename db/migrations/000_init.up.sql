CREATE SCHEMA IF NOT EXISTS db_local;

CREATE TABLE IF NOT EXISTS db_local.users(
    id serial PRIMARY KEY,
    email text NOT NULL,
    password text NOT NULL,
    name text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);