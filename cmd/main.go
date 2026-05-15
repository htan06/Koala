package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"koala.com/configs"
	"koala.com/internal/auth"
	"koala.com/internal/handler"
	"koala.com/internal/utils"
	"koala.com/middleware"
)

func main() {
	//Logger
	utils.Logger, _ = zap.NewDevelopment()

	//Jwt config
	jwtConfig := configs.GetJwtConfig()
	postgreCfg := configs.GetPostgreConfig()

	err := godotenv.Load()
	if err != nil {
		utils.Logger.Fatal("Load env failure")
	}

	Db, err := sqlx.Open("pgx", postgreCfg.DbUrl)

	if err != nil {
		utils.Logger.Fatal(err.Error())
	}

	userRepository := auth.NewUserRepository(Db)
	
	jwtService := auth.NewJwtService(*jwtConfig)
	
	authService := auth.NewAuthService(userRepository, jwtService)
	
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/login", authHandler.LoginHandle)

	r.POST("/change-password", middleware.JwtMiddleWare(jwtService), authHandler.ChangePasswordHandle)
	r.Run()
}
