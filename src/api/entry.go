package api

import (
	"database/sql"
	"net/http"

	db "github.com/bongster/golang_20210228/db/sqlc"
	"github.com/bongster/golang_20210228/token"
	"github.com/gin-gonic/gin"
)

// swagger:parameters getEntry
type getEntryRequest struct {
	// in: path
	// required: true
	ID int32 `json:"id" uri:"id" binding:"required,min=1"`
}

// swagger:response getEntryResponse
type getEntryResponse struct {
	// in: body
	Body db.Entry
}

// 	swagger:route GET /entries/{id} entries getEntry
// 		Responses:
// 			200: getEntryResponse
// 			400: badRequestResponse
func (server *Server) getEntry(ctx *gin.Context) {
	var req getEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

// listEntyResponse return response type from /entries API
// swagger:response listEntyResponse
type listEntyResponse struct {
	// in: body
	Body []db.Entry
}

// listEntries return user entry list
// 	swagger:route GET /entries entries listEntry
// 		Responses:
// 			200: listEntyResponse
// 			400: badRequestResponse
func (server *Server) listEntries(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByName(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	entries, err := server.store.ListEntries(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entries)
}
