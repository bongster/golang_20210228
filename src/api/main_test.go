package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/bongster/golang_20210228/token"
	"github.com/bongster/golang_20210228/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

var (
	testDB          *sql.DB
	testRouter      *gin.Engine
	store           *db.Store
	config          util.Config
	authoizationKey string = "Authorization"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func createNewUser(t *testing.T) db.User {
	hash, _ := util.HashPassword("password")
	arg := db.CreateUserParams{
		Username: fmt.Sprintf("daniel_%s", util.RandomString(10)),
		Password: hash,
		Balance:  1000,
		Currency: "SGD",
	}

	user, err := store.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Username, user.Username)
	return user
}
func createNewToken(t *testing.T) (db.User, string) {
	user := createNewUser(t)
	maker := token.NewJWTMaker(config.SecretKey)
	generatedToken, err := maker.CreateToken(user.Username, 15*time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, generatedToken)
	return user, generatedToken
}
func TestMain(m *testing.M) {
	var err error
	config, _ = util.LoadConfig("..", "test")
	fmt.Printf("%v", config)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}
	store = db.NewStore(testDB)
	testServer := NewServer(store, config)
	testRouter = testServer.router

	os.Exit(m.Run())
}
