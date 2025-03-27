package middleware

import (
	"net/http"
	"strings"

	"github.com/esuEdu/reurb-backend/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token format invalid"})
			ctx.Abort()
			return
		}

		token, err := util.ValidateToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid or expired"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Error processing token"})
			ctx.Abort()
			return
		}

		userID := claims["user_id"]

		ctx.Set("user_id", userID)

		ctx.Next()
	}
}
