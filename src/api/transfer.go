package api

import (
	"database/sql"
	"net/http"

	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/gin-gonic/gin"
)

// createTransferRequest user request parameter
type createTransferRequest struct {
	FromUserID int32 `json:"from_user_id" binding:"required,min=1"`
	ToUserID   int32 `json:"to_user_id" binding:"required,min=1"`
	Amount     int64 `json:"amount" binding:"required,gt=1"`
}

// ReqTransfer swagger parameters used in createTransfer
// swagger:parameters createTransfer
type ReqTransfer struct {
	// in: body
	// required: true
	Body createTransferRequest
}

// createTransferResponse response structure after success create user
// swagger:response createTransferResponse
type createTransferResponse struct {
	// in: body
	Body db.TransferTXResult
}

// 	swagger:route POST /transfers transfers createTransfer
// 		Responses:
// 			200: createTransferResponse
// 			400: badRequestResponse
func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check From user are whether exist or not
	if !server.validUser(ctx, req.FromUserID) {
		return
	}

	// Check To user are whether exist or not
	if !server.validUser(ctx, req.ToUserID) {
		return
	}

	arg := db.TransferTxParams{
		FromUserID: req.FromUserID,
		ToUserID:   req.ToUserID,
		Amount:     req.Amount,
	}
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) validUser(ctx *gin.Context, userID int32) bool {
	_, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	return true
}
