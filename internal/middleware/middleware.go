package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"koala.com/internal/auth"
	"koala.com/internal/utils"
)

func BaseMiddleWare() *gin.Engine {
	return gin.Default()
}

func JwtMiddleWare(jwtService auth.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		parts := strings.Split(authHeader, " ")
		if parts[0] == "Bearer" && len(parts) == 2 {
			tokenString := parts[1]

			claims, err := jwtService.ParseAccessToken(tokenString)

			if err != nil {
				c.Abort()
			}

			userIdString, err := claims.GetSubject()

			if err != nil {
				utils.Logger.Debug("Jwt middleware: " + err.Error())
				c.Abort()
			}

			userId, err := uuid.Parse(userIdString)

			if err != nil {
				utils.Logger.Debug("Jwt middleware: " + err.Error())
				c.Abort()
			}

			c.Set("userId", userId)
		}
		c.Next()
	}
}
