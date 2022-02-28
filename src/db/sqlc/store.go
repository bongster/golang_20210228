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
		// Checking user balance
		user1, err := store.GetUser(ctx, arg.FromUserID)
		if err != nil {
			return fmt.Errorf("from user [%d] is not exists", user1.ID)
		}
		user2, err := store.GetUser(ctx, arg.ToUserID)
		if err != nil {
			return fmt.Errorf("to user [%d] is not exists", user2.ID)
		}
		if user1.Balance < arg.Amount {
			return fmt.Errorf("user [%d] balance are not enough to send", user1.ID)
		}
		result.Transfer, err = q.CreatTransfer(ctx, CreatTransferParams{
			FromUserID: arg.FromUserID,
			ToUserID:   arg.ToUserID,
			Amount:     arg.Amount,
			Status:     "DONE", // TODO: update if the transfer api are running asynchonized
		})
		if err != nil {
			return err
		}
		// Create entry by user
		_, err = store.CreateEntry(ctx, CreateEntryParams{
			UserID:   user1.ID,
			Amount:   -arg.Amount,
			Currency: user1.Currency,
		})
		if err != nil {
			return err
		}
		// Create entry by user
		_, err = store.CreateEntry(ctx, CreateEntryParams{
			UserID:   user2.ID,
			Amount:   arg.Amount,
			Currency: user2.Currency,
		})
		if err != nil {
			return err
		}
		// Update accounts balance
		user1, err = store.UpdateUserBalance(ctx, UpdateUserBalanceParams{
			ID:      user1.ID,
			Balance: -arg.Amount,
		})
		if err != nil {
			return err
		}
		user2, err = store.UpdateUserBalance(ctx, UpdateUserBalanceParams{
			ID:      user2.ID,
			Balance: arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromUser = user1
		result.ToUser = user2
		return nil
	})
	return result, err
}
