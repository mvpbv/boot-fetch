-- name: CreateDuel :one
INSERT INTO duels (
    id,
    name,
    racer_1_id,
    racer_2_id,
    race_xp, 
    start_time
)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;
-- name: GetDuels :many
SELECT * FROM duels
WHERE completed = FALSE;
