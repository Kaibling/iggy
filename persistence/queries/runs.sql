-- name: FetchRunByWorkflow :many
SELECT runs.*, workflows.name as workflow_name, users.username as user_name
FROM runs
JOIN workflows
ON workflows.id = runs.id
JOIN users
ON users.id = runs.user_id
WHERE runs.workflow_id = $1;


-- name: FetchRuns :many
SELECT runs.*, workflows.name as workflow_name, users.username as user_name
FROM runs
JOIN workflows
ON workflows.id = runs.workflow_id
JOIN users
ON users.id = runs.user_id
WHERE runs.id = ANY($1::text[])
ORDER BY runs.id;


-- name: SaveRun :one
INSERT INTO runs (
  id,request_id,user_id, workflow_id, error, start_time, finish_time, created_at, modified_at, created_by, modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9,$10, $11
)
RETURNING id;

-- name: FetchRunByRequestID :one
SELECT runs.*, workflows.name as workflow_name, users.username as user_name
FROM runs
JOIN workflows
ON workflows.id = runs.workflow_id
JOIN users
ON users.id = runs.user_id
WHERE runs.request_id = $1
LIMIT 1;