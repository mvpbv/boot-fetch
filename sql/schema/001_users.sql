-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    boot_name TEXT NOT NULL UNIQUE,
    discord_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    wizard BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose Down
DROP TABLE users;