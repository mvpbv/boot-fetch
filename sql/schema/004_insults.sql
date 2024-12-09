-- +goose Up
CREATE TABLE insults (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    insult TEXT NOT NULL
);
-- +goose Down
DROP TABLE insults;