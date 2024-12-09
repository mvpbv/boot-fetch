-- name: CreateInsult :one
INSERT INTO insults (
    id,
    user_id,
    insult
)
VALUES(
    $1,
    $2,
    $3
)
RETURNING *;
-- name: GetUserInsults :many
SELECT * FROM insults WHERE user_id = $1;
