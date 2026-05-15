package main

import (
	"time"
	"koala.com/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"koala.com/config"
	"koala.com/handler"
	"koala.com/internal/auth"
	"koala.com/middleware"
)

func main() {
	utils.Logger, _ = zap.NewDevelopment()

	//Jwt config
	jwtConfig := config.JwtConfig{
		SecrectKeyAccess: []byte("dsdadadccececcsacscsa"),
		SecrectKeyRefresh: []byte("gsncicbds8tbdsahdbsa"),
		ExpAccessToken: time.Second * 600,
		ExpRefreshToken: time.Second * 86400,
	}

	Db := config.ConnectionDb()
	defer Db.Close()

	userRepository := auth.NewUserRepository(Db)
	
	jwtService := auth.NewJwtService(jwtConfig)
	
	authService := auth.NewAuthService(userRepository, jwtService)
	
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/login", authHandler.LoginHandle)

	r.POST("/change-password", middleware.JwtMiddleWare(jwtService), authHandler.ChangePasswordHandle)
	r.Run()
}
