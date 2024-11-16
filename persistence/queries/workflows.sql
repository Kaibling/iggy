-- name: FetchWorkflow :one
SELECT * FROM workflows
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: FetchAllWorkflows :many
SELECT * FROM workflows 
WHERE deleted_at IS NULL
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


-- name: UpdateWorkflow :exec
UPDATE workflows
SET
    name = COALESCE(sqlc.narg(name), name),
    code = COALESCE(sqlc.narg(code), code),
    object_type = COALESCE(sqlc.narg(object_type), object_type),
    fail_on_error = COALESCE(sqlc.narg(fail_on_error), fail_on_error),
    modified_at = $2,
    modified_by = $3
WHERE id = $1;