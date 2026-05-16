package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"koala.com/configs"
	"koala.com/internal/auth"
	"koala.com/internal/middleware"
	"koala.com/internal/rider"
	"koala.com/internal/utils"

	_ "github.com/jmoiron/sqlx"
)

func main() {
	//Load .env
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Warn("Load .env failure, using env default")
	}

	//Logger
	utils.Logger, _ = zap.NewDevelopment()

	//Jwt config
	jwtConfig := configs.GetJwtConfig()
	postgreCfg := configs.GetPostgreConfig()

	db, err := sqlx.Open("pgx", postgreCfg.DbUrl)

	if err != nil {
		utils.Logger.Fatal(err.Error())
	}

	//Repository
	userRepository := auth.NewUserRepository(db)
	riderRepository := rider.NewRiderRepository(db)

	//Service
	jwtService := auth.NewJwtService(jwtConfig)
	authService := auth.NewAuthService(userRepository, jwtService)
	riderService := rider.NewRiderService(riderRepository)

	//Handler
	authHandler := auth.NewAuthHandler(authService)
	riderHandler := rider.NewRiderHandler(riderService)

	r := gin.Default()

	r.GET("/rider/profile", middleware.JwtMiddleWare(jwtService), riderHandler.HandleGetProfile)
	r.POST("/auth/login", authHandler.HandleLogin)
	r.POST("/auth/change-password", middleware.JwtMiddleWare(jwtService), authHandler.HandleChangePassword)
	r.Run()
}
