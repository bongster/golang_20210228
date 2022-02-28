package db

import (
	"context"
	"testing"

	"github.com/bongster/golang_20210228/util"
	"github.com/stretchr/testify/require"
)

func createNewEntry(t *testing.T) Entry {
	user := createNewUser(t)
	arg := CreateEntryParams{
		UserID:   user.ID,
		Amount:   util.RandomInt(100, 1000),
		Currency: user.Currency,
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotZero(t, entry.ID)
	require.Equal(t, entry.UserID, user.ID)
	require.Equal(t, entry.Currency, user.Currency)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createNewEntry(t)
}

func TestListEntry(t *testing.T) {
	_, err := testQueries.ListEntries(context.Background(), 0)
	require.NoError(t, err)
}

func TestGetEntry(t *testing.T) {
	entry := createNewEntry(t)

	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.UserID, entry1.UserID)
	require.Equal(t, entry.Currency, entry1.Currency)
	require.Equal(t, entry.Amount, entry1.Amount)
}

func TestListEntryByUserId(t *testing.T) {
	entry := createNewEntry(t)
	entry2 := createNewEntry(t)
	entries, err := testQueries.ListEntries(context.Background(), entry.UserID)
	require.NoError(t, err)
	require.Equal(t, len(entries), 1)
	for _, e := range entries {
		require.NotEmpty(t, e)
		require.NotEqual(t, e.UserID, entry2.UserID)
		require.Equal(t, e.UserID, entry.UserID)
	}
}
