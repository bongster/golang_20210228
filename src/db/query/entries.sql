-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE
    (@user_id::integer = 0 OR user_id = @user_id)
ORDER BY id;

-- name: CreateEntry :one
INSERT INTO entries (
    user_id,
    amount,
    currency
) VALUES (
    $1, $2, $3
) RETURNING *;
