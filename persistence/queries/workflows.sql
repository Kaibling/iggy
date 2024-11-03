-- name: FetchWorkflow :one
SELECT * FROM workflows
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: FetchAllWorkflows :many
SELECT * FROM workflows 
ORDER BY id;

-- name: SaveWorkflow :one
INSERT INTO workflows (
  id,  name, code, object_type,fail_on_error,build_in, created_at, modified_at, created_by, modified_by, deleted_at 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id;

-- name: DeleteWorkflow :exec
UPDATE workflows
SET deleted_at = $1
WHERE id = $2;
