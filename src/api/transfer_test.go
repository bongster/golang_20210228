package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransferUnAuthoizedError(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/transfers", nil)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
func TestCreateTransfer(t *testing.T) {
	account1, token := createNewToken(t)
	account2 := createNewUser(t)
	w := httptest.NewRecorder()
	value := createTransferRequest{
		FromUserID: account1.ID,
		ToUserID:   account2.ID,
		Amount:     100,
	}
	jsonValue, _ := json.Marshal(value)
	req, _ := http.NewRequest("POST", "/transfers", bytes.NewBuffer(jsonValue))
	req.Header.Set(authoizationKey, token)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}
