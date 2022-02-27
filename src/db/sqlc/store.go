package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store Provider all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams contains the input parameter of the transfer tranasaction
type TransferTxParams struct {
	FromUserID int32 `json:"from_user_id"`
	ToUserID   int32 `json:"to_user_id"`
	Amount     int64 `json:"amount"`
}

// TransferTXResult is the result of the transfer transaction
type TransferTXResult struct {
	Transfer Transfer `json:"transfer"`
	FromUser User     `json:"from_user"`
	ToUser   User     `json:"to_user"`
}

// TransferTx performs a money transfer from one account to other account
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTXResult, error) {
	var result TransferTXResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreatTransfer(ctx, CreatTransferParams{
			FromUserID: arg.FromUserID,
			ToUserID:   arg.FromUserID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}
		// Update accounts balance
		return nil
	})
	return result, err
}
