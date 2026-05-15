package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"koala.com/configs"
	"koala.com/internal/auth"
	"koala.com/internal/handler"
	"koala.com/internal/utils"
	"koala.com/middleware"
)

func main() {
	utils.Logger, _ = zap.NewDevelopment()

	//Jwt config
	jwtConfig := configs.JwtConfig{
		SecrectKeyAccess: []byte("dsdadadccececcsacscsa"),
		SecrectKeyRefresh: []byte("gsncicbds8tbdsahdbsa"),
		ExpAccessToken: time.Second * 600,
		ExpRefreshToken: time.Second * 86400,
	}

	Db := configs.ConnectionDb()
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
