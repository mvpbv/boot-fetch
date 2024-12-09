-- +goose Up
ALTER TABLE users 
ADD COLUMN nickname TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE users
DROP COLUMN nickname;
