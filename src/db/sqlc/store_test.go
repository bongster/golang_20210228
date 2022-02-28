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
	ctx := context.Background()
	entryCount1, _ := store.GetCountEntry(context.Background(), account1.ID)
	entryCount2, _ := store.GetCountEntry(context.Background(), account2.ID)
	n := 5
	amount := int64(100)
	results := make(chan TransferTXResult)
	errs := make(chan error)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(ctx, TransferTxParams{
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

		_, err = store.GetTransfer(ctx, transfer.ID)
		require.NoError(t, err)
	}

	updatedAccount1, _ := store.GetUser(ctx, account1.ID)
	updatedAccount2, _ := store.GetUser(ctx, account2.ID)
	updatedEntryCount1, _ := store.GetCountEntry(ctx, account1.ID)
	updatedEntryCount2, _ := store.GetCountEntry(ctx, account2.ID)
	require.Equal(t, updatedAccount1.Balance-int64(n)*amount, account1.Balance)
	require.Equal(t, updatedAccount2.Balance+int64(n)*amount, account2.Balance)
	require.Equal(t, entryCount1+int64(n), updatedEntryCount1)
	require.Equal(t, entryCount2+int64(n), updatedEntryCount2)
}
