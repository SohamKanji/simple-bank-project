package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SohamKanji/simple-bank-project/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "Bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(maker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader(authorizationHeaderKey)

		if len(tokenString) == 0 {
			ctx.JSON(http.StatusUnauthorized, fmt.Errorf("authorization header is not provided"))
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(tokenString, " ")

		if len(tokenParts) != 2 || tokenParts[0] != authorizationTypeBearer {
			ctx.JSON(http.StatusUnauthorized, fmt.Errorf("authorization header is not valid"))
			ctx.Abort()
			return
		}

		token := tokenParts[1]

		payload, err := maker.VerifyToken(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, fmt.Errorf("invalid authentication token"))
			ctx.Abort()
			return
		}

		ctx.Set(authorizationPayloadKey, payload)

		ctx.Next()

	}
}
