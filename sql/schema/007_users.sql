-- +goose Up
ALTER TABLE users
ADD COLUMN target BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users
DROP COLUMN target;