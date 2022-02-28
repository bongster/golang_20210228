package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/bongster/golang_20210228/util"
	"github.com/stretchr/testify/require"
)

func TestRegistUser(t *testing.T) {
	w := httptest.NewRecorder()
	value := map[string]string{
		"username": fmt.Sprintf("user_%s", util.RandomString(10)),
		"password": util.RandomString(10),
		"currency": "SGD",
	}
	jsonValue, _ := json.Marshal(value)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	testRouter.ServeHTTP(w, req)
	require.Equal(t, 201, w.Code)
	var user db.User
	var err error
	json.Unmarshal(w.Body.Bytes(), &user)
	require.NotZero(t, user.ID)
	require.Equal(t, user.Username, value["username"])
	require.Equal(t, user.Currency, value["currency"])
	require.NotEqual(t, user.Password, value["password"])

	user, err = store.GetUserByName(context.Background(), value["username"])
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.Equal(t, value["username"], user.Username)
}
