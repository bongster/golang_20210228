package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/bongster/golang_20210228/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	var config util.Config
	config, err = util.LoadConfig("../..")
	fmt.Printf("%v", config)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
