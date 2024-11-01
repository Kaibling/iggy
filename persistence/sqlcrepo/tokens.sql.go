// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tokens.sql

package sqlcrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
  id, user_id,active,value,expires, created_at, created_by,updated_at,updated_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8,$9
)
RETURNING id, value, active, expires, created_at, created_by, updated_at, updated_by, user_id
`

type CreateTokenParams struct {
	ID        string
	UserID    string
	Active    int32
	Value     string
	Expires   pgtype.Timestamp
	CreatedAt pgtype.Timestamp
	CreatedBy string
	UpdatedAt pgtype.Timestamp
	UpdatedBy string
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRow(ctx, createToken,
		arg.ID,
		arg.UserID,
		arg.Active,
		arg.Value,
		arg.Expires,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Active,
		&i.Expires,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.UserID,
	)
	return i, err
}

const deleteTokenByValue = `-- name: DeleteTokenByValue :exec
DELETE FROM tokens
WHERE value = $1
`

func (q *Queries) DeleteTokenByValue(ctx context.Context, value string) error {
	_, err := q.db.Exec(ctx, deleteTokenByValue, value)
	return err
}

const getToken = `-- name: GetToken :one
SELECT id, value, active, expires, created_at, created_by, updated_at, updated_by, user_id FROM tokens
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, id string) (Token, error) {
	row := q.db.QueryRow(ctx, getToken, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Active,
		&i.Expires,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.UserID,
	)
	return i, err
}

const getTokenByValue = `-- name: GetTokenByValue :one
SELECT id, value, active, expires, created_at, created_by, updated_at, updated_by, user_id FROM tokens
WHERE value = $1 LIMIT 1
`

func (q *Queries) GetTokenByValue(ctx context.Context, value string) (Token, error) {
	row := q.db.QueryRow(ctx, getTokenByValue, value)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Active,
		&i.Expires,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.UserID,
	)
	return i, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, value, active, expires, created_at, created_by, updated_at, updated_by, user_id FROM tokens 
ORDER BY id
`

func (q *Queries) ListTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.Query(ctx, listTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Value,
			&i.Active,
			&i.Expires,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.UserID,
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

const listUserTokens = `-- name: ListUserTokens :many
SELECT id, value, active, expires, created_at, created_by, updated_at, updated_by, user_id FROM tokens 
WHERE user_id = $1
ORDER BY id
`

func (q *Queries) ListUserTokens(ctx context.Context, userID string) ([]Token, error) {
	rows, err := q.db.Query(ctx, listUserTokens, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Value,
			&i.Active,
			&i.Expires,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.UserID,
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

const listUserTokensByName = `-- name: ListUserTokensByName :many
SELECT tokens.id, tokens.value, tokens.active, tokens.expires, tokens.created_at, tokens.created_by, tokens.updated_at, tokens.updated_by, tokens.user_id FROM tokens 
join users on tokens.user_id = users.id
WHERE users.username = $1
ORDER BY tokens.id
`

func (q *Queries) ListUserTokensByName(ctx context.Context, username string) ([]Token, error) {
	rows, err := q.db.Query(ctx, listUserTokensByName, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Token
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.Value,
			&i.Active,
			&i.Expires,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.UserID,
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

const updateTokenValidity = `-- name: UpdateTokenValidity :one
UPDATE tokens
  set active = $2,
  updated_at = $3,
  updated_by = $4
WHERE id = $1
RETURNING id, value, active, expires, created_at, created_by, updated_at, updated_by, user_id
`

type UpdateTokenValidityParams struct {
	ID        string
	Active    int32
	UpdatedAt pgtype.Timestamp
	UpdatedBy string
}

func (q *Queries) UpdateTokenValidity(ctx context.Context, arg UpdateTokenValidityParams) (Token, error) {
	row := q.db.QueryRow(ctx, updateTokenValidity,
		arg.ID,
		arg.Active,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Active,
		&i.Expires,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.UserID,
	)
	return i, err
}
