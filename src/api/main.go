// Package api classification Payment API.
//
// Documentation for Payment API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Security:
// - authorization
// SecurityDefinitions:
// authorization:
//   type: apiKey
//   name: KEY
//   in: header
// swagger:meta
package api

import (
	"database/sql"
	"log"
	"os"

	db "github.com/bongster/golang_20210228/db/sqlc"
	_ "github.com/lib/pq"
)

// ValidationError use this error when invalid data comming
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// swagger:response genericError
type genericError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		Error error
	}
}

// swagger:response someResponse
type someResponse struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
	}
}

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:postgres@db:5432/payments?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

// Run HTTP api server
func Run() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	store := db.NewStore(conn)
	server := NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("can not start server:", err)
	}
}
