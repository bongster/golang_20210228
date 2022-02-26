-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransferByUserId :one
SELECT * FROM transfers
WHERE user_id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id;

-- name: DeletTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: CreatTransfer :one
INSERT INTO transfers (
    user_id,
    transfer_type,
    status,
    amount
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateTransferstatus :exec
UPDATE transfers SET status = $2, updated_at = now()
WHERE id = $1;