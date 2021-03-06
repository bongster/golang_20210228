-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    balance,
    currency
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateUserBalance :one
UPDATE users SET balance = balance + $2, updated_at = now()
WHERE id = $1
RETURNING *;