-- name: FetchUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: FetchUserByName :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: FetchAll :many
SELECT * FROM users 
ORDER BY id;

-- name: SaveUser :one
INSERT INTO users (
  id, username,password,active, created_at, created_by,modified_at,modified_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7,$8
)
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
  set password = $2,
  modified_at = $3,
  modified_by = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
