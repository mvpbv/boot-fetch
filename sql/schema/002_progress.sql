-- +goose Up

CREATE TABLE progress (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    time TIMESTAMP NOT NULL,
    level INTEGER NOT NULL,
    xp INTEGER NOT NULL,
    total_xp INTEGER NOT NULL,
    lessons INTEGER NOT NULL
    
);
-- +goose Down
DROP TABLE progress;