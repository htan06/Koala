package configs

import (
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
	"koala.com/internal/utils"
)

type PostgreConfig struct {
	DbUrl             string
	MaxOpenConnection int32
	MaxIdleConnection int32
	SslMode           string
}

func GetPostgreConfig() *PostgreConfig {
	maxOpenConnection, errMaxOpenConnection := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	maxIdleConnection, errMaxIdleConnection := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))

	if errMaxOpenConnection != nil {
		utils.Logger.Warn("Error pasering env DB_MAX_OPEN_CONNS, use default 10")
		maxOpenConnection = 10
	}

	if errMaxIdleConnection != nil {
		utils.Logger.Warn("Error pasering env DB_MAX_IDLE_CONNS, use default 5")
		maxIdleConnection = 5
	}

	return &PostgreConfig{
		DbUrl: os.Getenv("DATABASE_URL"),
		MaxOpenConnection: int32(maxOpenConnection),
		MaxIdleConnection: int32(maxIdleConnection),
		SslMode: os.Getenv("DB_SSLMODE"),
	}
}