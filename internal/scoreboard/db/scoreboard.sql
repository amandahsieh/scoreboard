-- name: GetAllScoreboards :many
SELECT * FROM scoreboards;

-- name: CreateScoreboard :one
INSERT INTO scoreboards (name)
VALUES ($1)
    RETURNING *;