package api

import (
	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
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

	r.POST("/users", server.createUser)
	r.GET("/users/:id", server.getUser)
	r.GET("/users", server.listUsers)

	r.POST("/transfers", server.createTransfer)

	// Set Doucumentation
	opts := middleware.RedocOpts{SpecURL: "./swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	r.GET("/docs", gin.WrapH(sh))
	r.StaticFile("/swagger.yml", "./swagger.yml")

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
