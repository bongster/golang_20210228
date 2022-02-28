package api

import (
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
