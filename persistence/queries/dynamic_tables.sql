-- name: FetchDynamicTables :many
SELECT * FROM dynamic_tables
WHERE id = ANY($1::text[])
ORDER BY id;

-- name: CreateDynamicTable :one
INSERT INTO dynamic_tables (
  id, table_name, created_at, created_by, modified_at, modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: DeleteDynamicTable :exec
DELETE FROM dynamic_tables
WHERE id = ANY($1::text[]);
