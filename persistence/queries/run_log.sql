-- name: FetchRunLogsByRun :many
SELECT * FROM run_logs
WHERE run_logs.run_id = $1;

-- name: FetchRunLog :one
SELECT * FROM run_logs
WHERE id = $1;

-- name: SaveRunLog :one
INSERT INTO run_logs (
  id, message, timestamp, run_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id;
