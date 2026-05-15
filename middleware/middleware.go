package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"koala.com/internal/auth"
)

func BaseMiddleWare() *gin.Engine {
	return gin.Default()
}

func JwtMiddleWare(jwtService *auth.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		parts := strings.Split(authHeader, " ")
		if parts[0] == "Bearer" && len(parts) == 2 {
			tokenString := parts[1]

			claims, err :=  jwtService.ValidAcessToken(tokenString)

			if err != nil {
				c.Abort()
			} else {
				username, err := claims.GetSubject()

				if err == nil {
					c.Set("username", username)
				}
			}
		}
		c.Next()
	}
}