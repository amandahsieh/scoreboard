-- name: GetAllScoreboards :many
SELECT * FROM scoreboards ORDER BY id;

-- name: CreateScoreboard :one
INSERT INTO scoreboards (name)
VALUES ($1)
    RETURNING *;

-- name: GetScoreboardByID :one
SELECT * FROM scoreboards WHERE id = $1;

-- name: UpdateScoreboard :one
UPDATE scoreboards
SET name = $2, updatedAt = NOW()
WHERE id = $1
    RETURNING *;

-- name: DeleteScoreboard :exec
DELETE FROM scoreboards WHERE id = $1;