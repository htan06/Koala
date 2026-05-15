package config

import "time"

type JwtConfig struct {
	SecrectKeyAccess  []byte
	SecrectKeyRefresh []byte
	ExpAccessToken    time.Duration
	ExpRefreshToken   time.Duration
}
