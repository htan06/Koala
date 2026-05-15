package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"koala.com/dto/auth/request"
	"koala.com/internal/auth"
	"koala.com/internal/utils"
)

type AuthHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (authHandler *AuthHandler) LoginHandle(c *gin.Context) {
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

func (authHandler *AuthHandler) ChangePasswordHandle(c *gin.Context) {
	var changePassword request.ChangePasswordDto
	ctx := c.Request.Context()

	errBindJson := c.ShouldBindJSON(&changePassword)
	if errBindJson != nil {
		utils.Logger.Debug(errBindJson.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	username, check := c.MustGet("username").(string)

	if !check {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}

	errChangePassword := authHandler.authService.ChangePassword(ctx, username, changePassword)

	if errChangePassword != nil {
		utils.Logger.Debug(username + " " + errChangePassword.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Change password failure",
		})
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Change password Success",
	})

}
