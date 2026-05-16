package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"koala.com/internal/dto/auth/request"
	"koala.com/internal/utils"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (authHandler *AuthHandler) HandleLogin(c *gin.Context) {
	var login request.LoginDto
	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&login)
	if err != nil {
		utils.Logger.Debug(err.Error())
		return
	}

	token, err := authHandler.authService.Login(ctx, login)

	if err != nil {
		utils.Logger.Debug(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

func (authHandler *AuthHandler) HandleChangePassword(c *gin.Context) {
	var changePassword request.ChangePasswordDto
	ctx := c.Request.Context()

	errBindJson := c.ShouldBindJSON(&changePassword)
	if errBindJson != nil {
		utils.Logger.Debug(errBindJson.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId := c.MustGet("userId").(uuid.UUID)

	errChangePassword := authHandler.authService.ChangePassword(ctx, userId, changePassword)

	if errChangePassword != nil {
		utils.Logger.Debug(userId.String() + " " + errChangePassword.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Change password failure",
		})
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Change password Success",
	})

}
