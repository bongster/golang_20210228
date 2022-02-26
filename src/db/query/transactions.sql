-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: GetTransactionByUserId :one
SELECT * FROM transactions
WHERE user_id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY id;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;

-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id,
    transaction_type,
    status,
    amount
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateTransactionStatus :exec
UPDATE transactions SET status = $2, updated_at = now()
WHERE id = $1;