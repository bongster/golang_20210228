package api

import (
	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/bongster/golang_20210228/token"
	"github.com/bongster/golang_20210228/util"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

// Server servers HTTP requests for our service
type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

const (
	secretKey = "asdadfsadfsadfsadfs12321asdas"
)

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store, config util.Config) *Server {
	maker := token.NewJWTMaker(config.SecretKey)
	server := &Server{
		store:      store,
		tokenMaker: maker,
	}
	router := gin.Default()
	router.POST("/login", server.loginUser)
	router.POST("/register", server.createUser)

	privateRouter := router.Group("/").Use(authMiddleware(maker))
	{
		privateRouter.GET("/users/:id", server.getUser)
		privateRouter.GET("/users", server.listUsers)
		privateRouter.POST("/transfers", server.createTransfer)
		privateRouter.GET("/transactions", server.listEntries)
		privateRouter.GET("/transactions/:id", server.getEntry)
	}

	// Set Doucumentation
	opts := middleware.RedocOpts{SpecURL: "./swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	router.GET("/docs", gin.WrapH(sh))
	router.StaticFile("/swagger.yml", "./swagger.yml")

	server.router = router
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
