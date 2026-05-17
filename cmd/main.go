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
	riderService := rider.NewRiderService(riderRepository, authService)

	//Handler
	authHandler := auth.NewAuthHandler(authService)
	riderHandler := rider.NewRiderHandler(riderService)

	//Middleware
	jwtMiddleWare := middleware.NewJwtMiddleware(jwtService)

	r := gin.Default()

	v1 := r.Group("api/v1")

	rider := v1.Group("/rider")
	rider.POST("/profile", jwtMiddleWare.Handler(), riderHandler.HanleAddProfile)
	rider.GET("/profile", jwtMiddleWare.Handler(), riderHandler.HandleGetProfile)
	rider.PATCH("/profile", jwtMiddleWare.Handler(), riderHandler.HanleUpadteProfile)
	rider.POST("/register", riderHandler.HadnleRegister)

	v1.POST("/auth/login", authHandler.HandleLogin)
	v1.POST("/auth/change-password", jwtMiddleWare.Handler(), authHandler.HandleChangePassword)
	
	r.Run("localhost:8080")
}
