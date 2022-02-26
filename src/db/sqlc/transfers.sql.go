// Code generated by sqlc. DO NOT EDIT.
// source: transfers.sql

package db

import (
	"context"
)

const creatTransfer = `-- name: CreatTransfer :one
INSERT INTO transfers (
    user_id,
    transfer_type,
    status,
    amount
) VALUES (
    $1, $2, $3, $4
) RETURNING id, user_id, transfer_type, status, amount, created_at, updated_at
`

type CreatTransferParams struct {
	UserID       int32  `json:"user_id"`
	TransferType string `json:"transfer_type"`
	Status       string `json:"status"`
	Amount       int64  `json:"amount"`
}

func (q *Queries) CreatTransfer(ctx context.Context, arg CreatTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, creatTransfer,
		arg.UserID,
		arg.TransferType,
		arg.Status,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransferType,
		&i.Status,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletTransfer = `-- name: DeletTransfer :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeletTransfer(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, user_id, transfer_type, status, amount, created_at, updated_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int32) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransferType,
		&i.Status,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTransferByUserId = `-- name: GetTransferByUserId :one
SELECT id, user_id, transfer_type, status, amount, created_at, updated_at FROM transfers
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetTransferByUserId(ctx context.Context, userID int32) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferByUserId, userID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransferType,
		&i.Status,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, user_id, transfer_type, status, amount, created_at, updated_at FROM transfers
ORDER BY id
`

func (q *Queries) ListTransfers(ctx context.Context) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TransferType,
			&i.Status,
			&i.Amount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransferstatus = `-- name: UpdateTransferstatus :exec
UPDATE transfers SET status = $2, updated_at = now()
WHERE id = $1
`

type UpdateTransferstatusParams struct {
	ID     int32  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateTransferstatus(ctx context.Context, arg UpdateTransferstatusParams) error {
	_, err := q.db.ExecContext(ctx, updateTransferstatus, arg.ID, arg.Status)
	return err
}