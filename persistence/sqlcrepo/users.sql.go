// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package sqlcrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const fetchAll = `-- name: FetchAll :many
SELECT id, username, password, active, created_at, created_by, modified_at, modified_by FROM users 
ORDER BY id
`

func (q *Queries) FetchAll(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, fetchAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Active,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.ModifiedAt,
			&i.ModifiedBy,
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

const fetchUser = `-- name: FetchUser :one
SELECT id, username, password, active, created_at, created_by, modified_at, modified_by FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) FetchUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, fetchUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Active,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}

const fetchUserByName = `-- name: FetchUserByName :one
SELECT id, username, password, active, created_at, created_by, modified_at, modified_by FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) FetchUserByName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, fetchUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Active,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}

const saveUser = `-- name: SaveUser :one
INSERT INTO users (
  id, username,password,active, created_at, created_by,modified_at,modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7,$8
)
RETURNING id, username, password, active, created_at, created_by, modified_at, modified_by
`

type SaveUserParams struct {
	ID         string
	Username   string
	Password   pgtype.Text
	Active     int32
	CreatedAt  pgtype.Timestamp
	CreatedBy  string
	ModifiedAt pgtype.Timestamp
	ModifiedBy string
}

func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) (User, error) {
	row := q.db.QueryRow(ctx, saveUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Active,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.ModifiedAt,
		arg.ModifiedBy,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Active,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE users
  set password = $2,
  modified_at = $3,
  modified_by = $4
WHERE id = $1
RETURNING id, username, password, active, created_at, created_by, modified_at, modified_by
`

type UpdatePasswordParams struct {
	ID         string
	Password   pgtype.Text
	ModifiedAt pgtype.Timestamp
	ModifiedBy string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRow(ctx, updatePassword,
		arg.ID,
		arg.Password,
		arg.ModifiedAt,
		arg.ModifiedBy,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Active,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}
