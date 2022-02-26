// Package classification Payment API.
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
// swagger:meta
package api

import "github.com/gin-gonic/gin"

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

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	meRouter := r.Group("/me")
	// 	swagger:route GET /me me me
	// 		Responses:
	// 			default: genericError
	// 			200: someResponse
	meRouter.GET("/", func(c *gin.Context) {
		me := make(map[string]interface{}, 0)
		c.JSON(200, me)
	})

	// swagger:route GET /me/balance me meBlanace
	// 		Responses:
	// 			default: genericError
	// 			200: someResponse
	meRouter.GET("/balance", func(c *gin.Context) {
		balance := make(map[string]interface{}, 0)
		c.JSON(200, balance)
	})

	// swagger:route GET /me/transactions me meTransactions
	// 		Responses:
	// 			default: genericError
	// 			200: someResponse
	meRouter.GET("/transactions", func(c *gin.Context) {
		transactions := make([]map[string]interface{}, 0)
		c.JSON(200, transactions)
	})

	// User API router
	userRouter := r.Group("/users")
	// swagger:route GET /users users listUsers
	// 		Responses:
	// 			default: genericError
	// 			200: someResponse
	userRouter.GET("/", func(c *gin.Context) {
		users := make([]map[string]interface{}, 0)
		c.JSON(200, users)
	})

	// Transaction router
	transactionRouter := r.Group("/transactions")
	// swagger:route GET /transactions transactions listTransactions
	// 		Responses:
	// 			default: genericError
	// 			200: someResponse
	transactionRouter.GET("/", func(c *gin.Context) {
		transactions := make([]map[string]interface{}, 0)
		c.JSON(200, transactions)
	})
	r.Run()
}
