-- name: FetchRun :one
SELECT runs.*,workflows.name FROM runs
JOIN workflows
ON workflows.id = runs.workflow_id
WHERE runs.id = $1
LIMIT 1;

-- name: FetchRunByWorkflow :many
SELECT runs.*,workflows.name FROM runs
JOIN workflows
ON workflows.id = runs.id
WHERE runs.workflow_id = $1;


-- name: SaveRun :one
INSERT INTO runs (
  id,request_id, workflow_id, error, start_time, finish_time, created_at, modified_at, created_by, modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9,$10
)
RETURNING id;
