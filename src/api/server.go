package api

import (
	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server servers HTTP requests for our service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	r := gin.Default()

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
	server.router = r
	return server
}

// Start Run HTTP server for our system
func (server *Server) Start(serverAddress string) error {
	return server.router.Run(serverAddress)
}
