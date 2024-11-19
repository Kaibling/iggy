// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: dynamic_table_variables.sql

package sqlcrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDynamicTableVariable = `-- name: CreateDynamicTableVariable :one
INSERT INTO dynamic_table_variables (
  id, name, variable_type, dynamic_table_id,  created_at, created_by, modified_at, modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id
`

type CreateDynamicTableVariableParams struct {
	ID             string
	Name           string
	VariableType   string
	DynamicTableID string
	CreatedAt      pgtype.Timestamp
	CreatedBy      string
	ModifiedAt     pgtype.Timestamp
	ModifiedBy     string
}

func (q *Queries) CreateDynamicTableVariable(ctx context.Context, arg CreateDynamicTableVariableParams) (string, error) {
	row := q.db.QueryRow(ctx, createDynamicTableVariable,
		arg.ID,
		arg.Name,
		arg.VariableType,
		arg.DynamicTableID,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.ModifiedAt,
		arg.ModifiedBy,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteDynamicTableVariables = `-- name: DeleteDynamicTableVariables :exec
DELETE
FROM dynamic_table_variables
WHERE id = ANY($1::text[])
`

func (q *Queries) DeleteDynamicTableVariables(ctx context.Context, dollar_1 []string) error {
	_, err := q.db.Exec(ctx, deleteDynamicTableVariables, dollar_1)
	return err
}

const fetchDynamicTableVariables = `-- name: FetchDynamicTableVariables :many
SELECT dynamic_table_variables.id, dynamic_table_variables.name, dynamic_table_variables.variable_type, dynamic_table_variables.dynamic_table_id, dynamic_table_variables.created_at, dynamic_table_variables.modified_at, dynamic_table_variables.created_by, dynamic_table_variables.modified_by,dynamic_tables.table_name as dynamic_table_name
FROM dynamic_table_variables
JOIN dynamic_tables
ON dynamic_tables.id = dynamic_table_variables.dynamic_table_id
WHERE dynamic_table_variables.id = ANY($1::text[])
ORDER BY dynamic_table_variables.id
`

type FetchDynamicTableVariablesRow struct {
	ID               string
	Name             string
	VariableType     string
	DynamicTableID   string
	CreatedAt        pgtype.Timestamp
	ModifiedAt       pgtype.Timestamp
	CreatedBy        string
	ModifiedBy       string
	DynamicTableName string
}

func (q *Queries) FetchDynamicTableVariables(ctx context.Context, dollar_1 []string) ([]FetchDynamicTableVariablesRow, error) {
	rows, err := q.db.Query(ctx, fetchDynamicTableVariables, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchDynamicTableVariablesRow
	for rows.Next() {
		var i FetchDynamicTableVariablesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.VariableType,
			&i.DynamicTableID,
			&i.CreatedAt,
			&i.ModifiedAt,
			&i.CreatedBy,
			&i.ModifiedBy,
			&i.DynamicTableName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchDynamicTableVariablesByDynamicTable = `-- name: FetchDynamicTableVariablesByDynamicTable :many
SELECT dynamic_table_variables.id, dynamic_table_variables.name, dynamic_table_variables.variable_type, dynamic_table_variables.dynamic_table_id, dynamic_table_variables.created_at, dynamic_table_variables.modified_at, dynamic_table_variables.created_by, dynamic_table_variables.modified_by,dynamic_tables.table_name as dynamic_table_name
FROM dynamic_table_variables
JOIN dynamic_tables
ON dynamic_tables.id = dynamic_table_variables.dynamic_table_id
WHERE dynamic_tables.id = $1
ORDER BY dynamic_table_variables.id
`

type FetchDynamicTableVariablesByDynamicTableRow struct {
	ID               string
	Name             string
	VariableType     string
	DynamicTableID   string
	CreatedAt        pgtype.Timestamp
	ModifiedAt       pgtype.Timestamp
	CreatedBy        string
	ModifiedBy       string
	DynamicTableName string
}

func (q *Queries) FetchDynamicTableVariablesByDynamicTable(ctx context.Context, id string) ([]FetchDynamicTableVariablesByDynamicTableRow, error) {
	rows, err := q.db.Query(ctx, fetchDynamicTableVariablesByDynamicTable, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchDynamicTableVariablesByDynamicTableRow
	for rows.Next() {
		var i FetchDynamicTableVariablesByDynamicTableRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.VariableType,
			&i.DynamicTableID,
			&i.CreatedAt,
			&i.ModifiedAt,
			&i.CreatedBy,
			&i.ModifiedBy,
			&i.DynamicTableName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
