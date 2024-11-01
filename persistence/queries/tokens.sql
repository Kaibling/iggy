-- name: GetToken :one
SELECT * FROM tokens
WHERE id = $1 LIMIT 1;

-- name: GetTokenByValue :one
SELECT * FROM tokens
WHERE value = $1 LIMIT 1;


-- name: ListUserTokens :many
SELECT * FROM tokens 
WHERE user_id = $1
ORDER BY id;

-- name: ListUserTokensByName :many
SELECT tokens.* FROM tokens 
join users on tokens.user_id = users.id
WHERE users.username = $1
ORDER BY tokens.id;

-- name: ListTokens :many
SELECT * FROM tokens 
ORDER BY id;

-- name: CreateToken :one
INSERT INTO tokens (
  id, user_id,active,value,expires, created_at, created_by,updated_at,updated_by
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8,$9
)
RETURNING *;

-- name: UpdateTokenValidity :one
UPDATE tokens
  set active = $2,
  updated_at = $3,
  updated_by = $4
WHERE id = $1
RETURNING *;


-- name: DeleteTokenByValue :exec
DELETE FROM tokens
WHERE value = $1;
