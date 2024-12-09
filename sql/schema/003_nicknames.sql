-- +goose Up

CREATE TABLE nicknames (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    nickname TEXT NOT NULL
);
-- +goose Down

DROP TABLE nicknames;