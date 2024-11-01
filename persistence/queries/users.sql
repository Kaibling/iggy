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
  id, username,password,active, created_at, created_by,updated_at,updated_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7,$8
)
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
  set password = $2,
  updated_at = $3,
  updated_by = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
