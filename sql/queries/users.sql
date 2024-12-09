-- name: CreateUser :one
INSERT INTO users (id, boot_name, discord_name, created_at, updated_at, wizard, nickname)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET updated_at = $2, nickname = $3 WHERE id = $1;


-- name: GetUserIdByDiscordName :one
SELECT id FROM users WHERE discord_name = $1;
-- name: GetUser :one
SELECT * FROM users WHERE id = $1;
-- name: MakeWizard :one
UPDATE users SET wizard = True WHERE id = $1 RETURNING *;
-- name: GetUsers :many
SELECT * FROM users;
-- name: GetWizards :many
SELECT * FROM users WHERE wizard = True;