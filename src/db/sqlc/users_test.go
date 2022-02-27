package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/bongster/golang_20210228/util"
	"github.com/stretchr/testify/require"
)

func createNewUser(t *testing.T) User {
	hash, _ := util.HashPassword("password")
	arg := CreateUserParams{
		Username: fmt.Sprintf("daniel_%s", util.RandomString(10)),
		Password: hash,
		Balance:  0,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Username, user.Username)
	return user
}
func TestCreateUser(t *testing.T) {
	createNewUser(t)
}

func TestListUser(t *testing.T) {
	_, err := testQueries.ListUsers(context.Background())
	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	user := createNewUser(t)
	u2, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, user.ID, u2.ID)
	require.Equal(t, user.Password, u2.Password)
	require.Equal(t, user.Username, u2.Username)
}
