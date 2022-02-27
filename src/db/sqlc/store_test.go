package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createNewUser(t)
	account2 := createNewUser(t)
	// run concurrent transfer transactions

	n := 5
	amount := int64(100)
	results := make(chan TransferTXResult)
	errs := make(chan error)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromUserID: account1.ID,
				ToUserID:   account2.ID,
				Amount:     amount,
			})
			results <- result
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		transfer := result.Transfer
		require.NotEmpty(t, transfer)

		require.Equal(t, account1.ID, transfer.FromUserID)
		require.Equal(t, account2.ID, transfer.ToUserID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// TODO: check user balance
	}
}
