package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"live-pilot/pkg/conf"
	"strconv"
	"time"
)

func GenerateAccessToken(id uint32) (string, error) {
	jwtConfig := conf.GetConfig().Jwt

	expiresAt := time.Now().Add(time.Minute * time.Duration(jwtConfig.ValidWithinMinutes))
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(id)),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        strconv.Itoa(int(id)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtConfig.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
