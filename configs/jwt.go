package configs

import (
	"os"
	"strconv"
	"time"

	"koala.com/internal/utils"
)

type JwtConfig struct {
	SecrectKeyAccess  []byte
	SecrectKeyRefresh []byte
	ExpAccessToken    time.Duration
	ExpRefreshToken   time.Duration
}

func GetJwtConfig() *JwtConfig {

	secretKeyAccess := []byte(os.Getenv("SECRET_KEY_ACCESS"))
	secretKeyRefresh := []byte(os.Getenv("SECRET_KEY_REFRESH"))
	expAccessToken, errExpAccess := strconv.Atoi(os.Getenv("EXP_ACCESS_TOKEN"))
	expRefreshToken, errExpRefresh := strconv.Atoi(os.Getenv("EXP_REFRESH_TOKEN"))

	if len(secretKeyAccess) == 0 {
		utils.Logger.Fatal("CRITICAL: SECRET_KEY_ACCESS is missing in environment variables. Application cannot start.")
	}

	if len(secretKeyRefresh) == 0 {
		utils.Logger.Fatal("CRITICAL: EXP_REFRESH_TOKEN is missing in environment variables. Application cannot start.")
	}

	if errExpAccess != nil {
		expAccessToken = 600
		utils.Logger.Warn("Err pasering exp access token, use default EXP_ACCESS_TOKEN: 600")
	}

	if errExpRefresh != nil {
		utils.Logger.Warn("Err pasering exp refresh token, use default EXP_REFRESH_TOKEN: 86400")
		expRefreshToken = 86400
	}

	return &JwtConfig{
		SecrectKeyAccess:  []byte(os.Getenv("SECRET_KEY_ACCESS")),
		SecrectKeyRefresh: []byte(os.Getenv("SECRET_KEY_REFRESH")),
		ExpAccessToken:    time.Second * time.Duration(expAccessToken),
		ExpRefreshToken:   time.Second * time.Duration(expRefreshToken),
	}
}
