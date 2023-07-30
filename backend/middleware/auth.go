package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	utils "github.com/bitFieldE/go-next/backend/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	token, err := utils.JwtTokenParser(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
