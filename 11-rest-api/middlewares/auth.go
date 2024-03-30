package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"res-api.com/apis/utils"
)

func Authenticated(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		context.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.Set("userId", userId)
	context.Next()
}
