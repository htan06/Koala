package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"koala.com/internal/auth"
	"koala.com/internal/utils"
)

type JwtMiddleware struct {
	jwtService auth.JwtService
}

func NewJwtMiddleware(jwtService auth.JwtService) *JwtMiddleware {
	return &JwtMiddleware{jwtService}
} 

func (jwtMiddleWare *JwtMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
			c.Abort()
			return
		} 
		
		tokenString := parts[1]

		claims, err := jwtMiddleWare.jwtService.ParseAccessToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
			c.Abort()
			return
		}
		
		
		userIdString, err := claims.GetSubject()

		if err != nil {
			utils.Logger.Debug("Jwt middleware: " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
			c.Abort()
			return
		}

		userId, err := uuid.Parse(userIdString)

		if err != nil {
			utils.Logger.Debug("Jwt middleware: " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorization"})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		
		c.Next()
	}
}
