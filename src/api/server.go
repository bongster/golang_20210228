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

	// swagger:route POST /users users createUser
	// 	Response:
	// 		default: genericError
	// 		200: somResponse
	r.POST("/users", server.createUser)

	server.router = r
	return server
}

// Start Run HTTP server for our system
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse generate error response
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
