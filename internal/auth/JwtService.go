package auth

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"koala.com/configs"
)

type JwtService struct {
	cfg configs.JwtConfig
}

func NewJwtService(jwtCfg configs.JwtConfig) *JwtService {
	return &JwtService{jwtCfg}
}

func (jwtService *JwtService) generateAccessToken(username string) string {
	claims := jwt.MapClaims{
		"iss": "koala",
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(jwtService.cfg.ExpAccessToken).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtService.cfg.SecrectKeyAccess)

	if err != nil {
		slog.Error("Err generate access token for %s", username)
		return ""
	} else {
		return tokenString
	}
}

func (jwtService *JwtService) generateRefreshToken(username string) string {
	claims := jwt.MapClaims{
		"iss": "koala",
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(jwtService.cfg.ExpRefreshToken).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtService.cfg.SecrectKeyRefresh)

	if err != nil {
		slog.Error("Err generate refresh token for %s", username)
		return ""
	} else {
		return tokenString
	}
}

func (jwtService *JwtService) ValidAcessToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtService.cfg.SecrectKeyAccess, nil
	})
	return token.Claims, err
}