-- name: GetUserProgress :many
SELECT progress.level, progress.xp, progress.lessons, progress.total_xp, progress.time, users.discord_name
FROM progress
JOIN users ON users.id = progress.user_id
WHERE users.id = $1
ORDER BY time DESC;

-- name: GetRecentProgressUser :one
SELECT progress.lessons, progress.time, progress.level, progress.xp, progress.total_xp
FROM progress
JOIN users ON users.id = progress.user_id
WHERE users.id = $1
ORDER BY time DESC
LIMIT 1;

-- name: GetFirstProgressUser :one
SELECT progress.lessons, progress.time, progress.level, progress.xp, progress.total_xp
FROM progress
JOIN users ON users.id = progress.user_id
WHERE users.id = $1
ORDER BY time ASC
LIMIT 1;
-- name: GetWeeklyStartProgressUser :one
SELECT progress.lessons, progress.time, progress.level, progress.xp, progress.total_xp
FROM progress
JOIN users ON users.id = progress.user_id
WHERE users.id = $1 AND progress.time <= $2
ORDER BY time DESC
LIMIT 1;
