
-- name: CreateUsers :one
INSERT INTO users (
  id, username,role
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUsers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- -- name: GetAccountForUpdate :one
-- SELECT * FROM accounts
-- WHERE id = $1 LIMIT 1
-- FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
WHERE Role = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateUsers :one
UPDATE users
SET Role = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE id = $1; 


