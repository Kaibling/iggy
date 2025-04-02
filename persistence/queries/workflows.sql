-- name: FetchWorkflow :one
SELECT * FROM workflows
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: FetchAllWorkflows :many
SELECT * FROM workflows 
WHERE deleted_at IS NULL
ORDER BY id;

-- name: FetchToBackup :many
SELECT id FROM 
workflows 
WHERE
deleted_at IS NULL AND build_in = false
ORDER BY id;

-- name: FetchBackupAll :many
SELECT id FROM 
workflows 
WHERE
deleted_at IS NULL
ORDER BY id;

-- name: SaveWorkflow :one
INSERT INTO workflows (
  id,  name, code, object_type,fail_on_error,build_in, created_at, modified_at, created_by, modified_by, deleted_at 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id;


-- name: UpsertWorkflow :exec
WITH upsert AS (
  UPDATE workflows
  SET 
    code = $3,
    object_type = $4,
    fail_on_error = $5,
    modified_at = $8,
    modified_by = $10
  WHERE id = $1 OR (name = $2 AND deleted_at = $11)
  RETURNING *
)
INSERT INTO workflows (
  id, name, code, object_type, fail_on_error, build_in, created_at, modified_at, created_by, modified_by, deleted_at
)
SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
WHERE NOT EXISTS (SELECT 1 FROM upsert);



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