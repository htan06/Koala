package auth

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"koala.com/configs"
)

type JwtService interface {
	generateAccessToken(userId uuid.UUID) (string, error)
	generateRefreshToken(userId uuid.UUID) (string, error)
	ParseAccessToken(tokenString string) (jwt.Claims, error)
	ParseRefreshToken(tokenString string) (jwt.Claims, error)
}

type JwtServiceImpl struct {
	cfg *configs.JwtConfig
}

func NewJwtService(jwtCfg *configs.JwtConfig) *JwtServiceImpl {
	return &JwtServiceImpl{jwtCfg}
}

func (jwtService *JwtServiceImpl) generateAccessToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"iss": "koala",
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(jwtService.cfg.ExpAccessToken).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtService.cfg.SecrectKeyAccess)

	if err != nil {
		return "", fmt.Errorf("Generate access token err: %w", err)
	} else {
		return tokenString, nil
	}
}

func (jwtService *JwtServiceImpl) generateRefreshToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"iss": "koala",
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(jwtService.cfg.ExpRefreshToken).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtService.cfg.SecrectKeyRefresh)

	if err != nil {
		return "", fmt.Errorf("Generate refresh token err: %w", err)
	} else {
		return tokenString, nil
	}
}

func (jwtService *JwtServiceImpl) ParseAccessToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtService.cfg.SecrectKeyAccess, nil
	})
	return token.Claims, err
}

func (jwtService *JwtServiceImpl) ParseRefreshToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtService.cfg.SecrectKeyRefresh, nil
	})
	return token.Claims, err
}