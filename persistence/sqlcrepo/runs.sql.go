// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: runs.sql

package sqlcrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const fetchRun = `-- name: FetchRun :one
SELECT runs.id, runs.request_id, runs.workflow_id, runs.error, runs.start_time, runs.finish_time, runs.created_at, runs.modified_at, runs.created_by, runs.modified_by,workflows.name FROM runs
JOIN workflows
ON workflows.id = runs.workflow_id
WHERE runs.id = $1
LIMIT 1
`

type FetchRunRow struct {
	ID         string
	RequestID  pgtype.Text
	WorkflowID string
	Error      pgtype.Text
	StartTime  pgtype.Timestamp
	FinishTime pgtype.Timestamp
	CreatedAt  pgtype.Timestamp
	ModifiedAt pgtype.Timestamp
	CreatedBy  string
	ModifiedBy string
	Name       string
}

func (q *Queries) FetchRun(ctx context.Context, id string) (FetchRunRow, error) {
	row := q.db.QueryRow(ctx, fetchRun, id)
	var i FetchRunRow
	err := row.Scan(
		&i.ID,
		&i.RequestID,
		&i.WorkflowID,
		&i.Error,
		&i.StartTime,
		&i.FinishTime,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.CreatedBy,
		&i.ModifiedBy,
		&i.Name,
	)
	return i, err
}

const fetchRunByWorkflow = `-- name: FetchRunByWorkflow :many
SELECT runs.id, runs.request_id, runs.workflow_id, runs.error, runs.start_time, runs.finish_time, runs.created_at, runs.modified_at, runs.created_by, runs.modified_by,workflows.name FROM runs
JOIN workflows
ON workflows.id = runs.id
WHERE runs.workflow_id = $1
`

type FetchRunByWorkflowRow struct {
	ID         string
	RequestID  pgtype.Text
	WorkflowID string
	Error      pgtype.Text
	StartTime  pgtype.Timestamp
	FinishTime pgtype.Timestamp
	CreatedAt  pgtype.Timestamp
	ModifiedAt pgtype.Timestamp
	CreatedBy  string
	ModifiedBy string
	Name       string
}

func (q *Queries) FetchRunByWorkflow(ctx context.Context, workflowID string) ([]FetchRunByWorkflowRow, error) {
	rows, err := q.db.Query(ctx, fetchRunByWorkflow, workflowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchRunByWorkflowRow
	for rows.Next() {
		var i FetchRunByWorkflowRow
		if err := rows.Scan(
			&i.ID,
			&i.RequestID,
			&i.WorkflowID,
			&i.Error,
			&i.StartTime,
			&i.FinishTime,
			&i.CreatedAt,
			&i.ModifiedAt,
			&i.CreatedBy,
			&i.ModifiedBy,
			&i.Name,
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

const saveRun = `-- name: SaveRun :one
INSERT INTO runs (
  id,request_id, workflow_id, error, start_time, finish_time, created_at, modified_at, created_by, modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9,$10
)
RETURNING id
`

type SaveRunParams struct {
	ID         string
	RequestID  pgtype.Text
	WorkflowID string
	Error      pgtype.Text
	StartTime  pgtype.Timestamp
	FinishTime pgtype.Timestamp
	CreatedAt  pgtype.Timestamp
	ModifiedAt pgtype.Timestamp
	CreatedBy  string
	ModifiedBy string
}

func (q *Queries) SaveRun(ctx context.Context, arg SaveRunParams) (string, error) {
	row := q.db.QueryRow(ctx, saveRun,
		arg.ID,
		arg.RequestID,
		arg.WorkflowID,
		arg.Error,
		arg.StartTime,
		arg.FinishTime,
		arg.CreatedAt,
		arg.ModifiedAt,
		arg.CreatedBy,
		arg.ModifiedBy,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
