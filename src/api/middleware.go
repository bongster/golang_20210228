package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bongster/golang_20210228/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationPayloadKey = "payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authroizationHeader := ctx.GetHeader("authorization")
		if len(authroizationHeader) == 0 {
			err := errors.New("authoization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		}
		fields := strings.Fields(authroizationHeader)
		if len(fields) > 2 {
			err := errors.New("invalid authoization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		}
		if strings.ToLower(fields[0]) != "bearer" {
			err := errors.New("unsupported authorization type")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		payload, err := tokenMaker.Verify(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, payload)
			return
		}
		ctx.Set("payload", payload)
		ctx.Next()
	}
}
