-- +goose Up
CREATE TABLE duels (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    racer_1_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    racer_2_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    race_xp INTEGER NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    winner_id UUID REFERENCES users(id) ON DELETE CASCADE,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP
);
-- +goose Down
DROP TABLE duels;