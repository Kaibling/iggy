-- name: FetchDynamicFields :many
SELECT dynamic_table_variables.*,dynamic_tables.table_name as dynamic_table_name
FROM dynamic_table_variables
JOIN dynamic_tables
ON dynamic_tables.id = dynamic_table_variables.dynamic_table_id
WHERE dynamic_table_variables.id = ANY($1::text[])
ORDER BY dynamic_table_variables.id;

-- name: FetchDynamicFieldsByDynamicTable :many
SELECT dynamic_table_variables.*,dynamic_tables.table_name as dynamic_table_name
FROM dynamic_table_variables
JOIN dynamic_tables
ON dynamic_tables.id = dynamic_table_variables.dynamic_table_id
WHERE dynamic_tables.id = $1
ORDER BY dynamic_table_variables.id;

-- name: FetchDynamicFieldsByDynamicTableName :many
SELECT dynamic_table_variables.*,dynamic_tables.table_name as dynamic_table_name
FROM dynamic_table_variables
JOIN dynamic_tables
ON dynamic_tables.id = dynamic_table_variables.dynamic_table_id
WHERE dynamic_tables.table_name = $1
ORDER BY dynamic_table_variables.id;



-- name: CreateDynamicField :one
INSERT INTO dynamic_table_variables (
  id, name, variable_type, dynamic_table_id,  created_at, created_by, modified_at, modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id;

-- name: DeleteDynamicFields :exec
DELETE
FROM dynamic_table_variables
WHERE id = ANY($1::text[]);
