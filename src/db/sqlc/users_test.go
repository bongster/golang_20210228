package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/bongster/golang_20210228/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	hash, _ := util.HashPassword("password")
	arg := CreateUserParams{
		Username: fmt.Sprintf("daniel_%s", util.RandomString(10)),
		Password: hash,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Username, user.Username)
}
