package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/bongster/golang_20210228/util"
	"github.com/gin-gonic/gin"
)

// createUserRequest user request parameter
type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=SGD"`
}

// Req swagger parameters used in createUser
// swagger:parameters createUser
type Req struct {
	// in: body
	// required: true
	Body createUserRequest
}

// createUserResponse response structure after success create user
// swagger:response createUserResponse
type createUserResponse struct {
	// in: body
	Body db.User
}

// badRequestResponse response structure after invalid input from body
// swagger:response badRequestResponse
type badRequestResponse struct {
	// in: body
	Body struct {
		error error
	}
}

// 	swagger:route POST /register users createUser
// 		Responses:
// 			200: createUserResponse
// 			400: badRequestResponse
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashPassword,
		Currency: req.Currency,
		Balance:  0,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

// swagger:parameters getUser
type getUserRequest struct {
	// in: path
	// required: true
	ID int32 `json:"id" uri:"id" binding:"required,min=1"`
}

// swagger:response getUserResponse
type getUserResponse struct {
	// in: body
	Body db.User
}

// 	swagger:route GET /users/{id} users getUser
// 		Responses:
// 			200: getUserResponse
// 			400: badRequestResponse
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// swagger:response listUserResponse
type listUserResponse struct {
	// in: body
	Body []db.User
}

// 	swagger:route GET /users users listUser
// 		Responses:
// 			200: listUserResponse
// 			400: badRequestResponse
func (server *Server) listUsers(ctx *gin.Context) {
	users, err := server.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// loginUserRequest user login request parameter
type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Req swagger parameters used in loginUser
// swagger:parameters loginUser
type loginUserReq struct {
	// in: body
	// required: true
	Body loginUserRequest
}

// swagger:response loginUserResponse
type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}

// 	swagger:route GET /login users loginUser
// 		Responses:
// 			200: loginUserResponse
// 			400: badRequestResponse
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUserByName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if err := util.CheckPasswordHash(req.Password, user.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	token, err := server.tokenMaker.CreateToken(user.Username, 15*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := loginUserResponse{
		AccessToken: token,
	}
	ctx.JSON(http.StatusOK, res)
}
