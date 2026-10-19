package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

func JWTMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		claims, err := ParseToken(strings.TrimSpace(parts[1]), secret)
		if err != nil {
			if errors.Is(err, ErrExpiredToken) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}

// GetUserID — хелпер для обработчиков, читает ID из контекста Gin.
func GetUserID(c *gin.Context) (uint, bool) {
	val, exists := c.Get(ContextUserIDKey)
	if !exists {
		return 0, false
	}
	id, ok := val.(uint)
	return id, ok
}