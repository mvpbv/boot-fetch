-- name: CreateProgress :one
INSERT INTO progress (
    id,
    user_id,
    time,
    level,
    xp,
    total_xp,
    lessons
)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;