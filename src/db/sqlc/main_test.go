package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/bongster/golang_20210228/migration"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"
	dbName   = "testdb"
)

var testQueries *Queries

func SetUpTestDB() {
	conninfo := "user=postgres password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}
	db.Exec("CREATE DATABASE " + dbName)
}
func TestMain(m *testing.M) {
	SetUpTestDB()
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}
	testQueries = New(conn)

	migration.Run(dbSource, "file://../../db/migration")

	os.Exit(m.Run())
}
