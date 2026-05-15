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

	expAccessToken, errExpAccess := strconv.Atoi(os.Getenv("EXP_ACCESS_TOKEN"))
	expRefreshToken, errExpRefresh := strconv.Atoi(os.Getenv("EXP_REFRESH_TOKEN"))

	if errExpAccess != nil {
		expAccessToken = 600
		utils.Logger.Warn("Err pasering exp access token, use default value 600")
	}
	
	if errExpRefresh != nil {
		utils.Logger.Warn("Err pasering exp refresh token, use default value 86400")
		expRefreshToken = 86400
	}

	return &JwtConfig{
		SecrectKeyAccess: []byte(os.Getenv("SECRET_KEY_ACCESS")),
		SecrectKeyRefresh: []byte(os.Getenv("SECRET_KEY_REFRESH")),
		ExpAccessToken: time.Second * time.Duration(expAccessToken),
		ExpRefreshToken: time.Second * time.Duration(expRefreshToken),
	}
}