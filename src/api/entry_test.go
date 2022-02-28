package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createNewEntry(t *testing.T, user db.User) db.Entry {
	args := db.CreateEntryParams{
		UserID:   user.ID,
		Amount:   100,
		Currency: user.Currency,
	}
	entry, err := store.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.UserID, user.ID)
	require.Equal(t, entry.Currency, user.Currency)
	return entry
}

func TestListEntry(t *testing.T) {
	user, token := createNewToken(t)
	createNewEntry(t, user)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/entries", nil)
	req.Header.Set(authoizationKey, token)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var body []db.Entry
	json.Unmarshal(w.Body.Bytes(), &body)
	require.NotZero(t, len(body))
	for _, v := range body {
		require.NotZero(t, v.UserID)
		require.Equal(t, v.UserID, user.ID)
		require.Equal(t, v.Currency, user.Currency)
	}
}

func TestGetEntry(t *testing.T) {
	user, token := createNewToken(t)
	entry := createNewEntry(t, user)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/entries/%d", entry.ID), nil)
	req.Header.Set(authoizationKey, token)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var body db.Entry
	json.Unmarshal(w.Body.Bytes(), &body)
	require.NotZero(t, body.ID)
	require.Equal(t, body.UserID, entry.UserID)
	require.Equal(t, body.Currency, entry.Currency)
	require.Equal(t, body.Amount, entry.Amount)
}

func TestGetEntryNotFound(t *testing.T) {
	_, token := createNewToken(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/entries/%d", 123123123), nil)
	req.Header.Set(authoizationKey, token)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

func TestGetEntryInValidUri(t *testing.T) {
	_, token := createNewToken(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/entries/%d", 0), nil)
	req.Header.Set(authoizationKey, token)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
