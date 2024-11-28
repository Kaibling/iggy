// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: workflows.sql

package sqlcrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteWorkflow = `-- name: DeleteWorkflow :exec
UPDATE workflows
SET deleted_at = $1
WHERE id = $2
`

type DeleteWorkflowParams struct {
	DeletedAt pgtype.Timestamp
	ID        string
}

func (q *Queries) DeleteWorkflow(ctx context.Context, arg DeleteWorkflowParams) error {
	_, err := q.db.Exec(ctx, deleteWorkflow, arg.DeletedAt, arg.ID)
	return err
}

const fetchAllWorkflows = `-- name: FetchAllWorkflows :many
SELECT id, name, code, object_type, fail_on_error, build_in, created_at, modified_at, created_by, modified_by, deleted_at FROM workflows 
WHERE deleted_at IS NULL
ORDER BY id
`

func (q *Queries) FetchAllWorkflows(ctx context.Context) ([]Workflow, error) {
	rows, err := q.db.Query(ctx, fetchAllWorkflows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Workflow
	for rows.Next() {
		var i Workflow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.ObjectType,
			&i.FailOnError,
			&i.BuildIn,
			&i.CreatedAt,
			&i.ModifiedAt,
			&i.CreatedBy,
			&i.ModifiedBy,
			&i.DeletedAt,
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

const fetchToBackup = `-- name: FetchToBackup :many
SELECT id FROM 
workflows 
WHERE
deleted_at IS NULL AND build_in = false
ORDER BY id
`

func (q *Queries) FetchToBackup(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, fetchToBackup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchWorkflow = `-- name: FetchWorkflow :one
SELECT id, name, code, object_type, fail_on_error, build_in, created_at, modified_at, created_by, modified_by, deleted_at FROM workflows
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) FetchWorkflow(ctx context.Context, id string) (Workflow, error) {
	row := q.db.QueryRow(ctx, fetchWorkflow, id)
	var i Workflow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.ObjectType,
		&i.FailOnError,
		&i.BuildIn,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.CreatedBy,
		&i.ModifiedBy,
		&i.DeletedAt,
	)
	return i, err
}

const saveWorkflow = `-- name: SaveWorkflow :one
INSERT INTO workflows (
  id,  name, code, object_type,fail_on_error,build_in, created_at, modified_at, created_by, modified_by, deleted_at 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id
`

type SaveWorkflowParams struct {
	ID          string
	Name        string
	Code        pgtype.Text
	ObjectType  string
	FailOnError bool
	BuildIn     bool
	CreatedAt   pgtype.Timestamp
	ModifiedAt  pgtype.Timestamp
	CreatedBy   string
	ModifiedBy  string
	DeletedAt   pgtype.Timestamp
}

func (q *Queries) SaveWorkflow(ctx context.Context, arg SaveWorkflowParams) (string, error) {
	row := q.db.QueryRow(ctx, saveWorkflow,
		arg.ID,
		arg.Name,
		arg.Code,
		arg.ObjectType,
		arg.FailOnError,
		arg.BuildIn,
		arg.CreatedAt,
		arg.ModifiedAt,
		arg.CreatedBy,
		arg.ModifiedBy,
		arg.DeletedAt,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const updateWorkflow = `-- name: UpdateWorkflow :exec
UPDATE workflows
SET
    name = COALESCE($4, name),
    code = COALESCE($5, code),
    object_type = COALESCE($6, object_type),
    fail_on_error = COALESCE($7, fail_on_error),
    modified_at = $2,
    modified_by = $3
WHERE id = $1
`

type UpdateWorkflowParams struct {
	ID          string
	ModifiedAt  pgtype.Timestamp
	ModifiedBy  string
	Name        pgtype.Text
	Code        pgtype.Text
	ObjectType  pgtype.Text
	FailOnError pgtype.Bool
}

func (q *Queries) UpdateWorkflow(ctx context.Context, arg UpdateWorkflowParams) error {
	_, err := q.db.Exec(ctx, updateWorkflow,
		arg.ID,
		arg.ModifiedAt,
		arg.ModifiedBy,
		arg.Name,
		arg.Code,
		arg.ObjectType,
		arg.FailOnError,
	)
	return err
}
