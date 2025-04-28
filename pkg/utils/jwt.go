package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"live-pilot/pkg/conf"
	"strconv"
	"time"
)

type JwtClaims struct {
	jwt.RegisteredClaims
}

func GenerateAccessToken(id uint32) (string, error) {
	jwtConfig := conf.GetConfig().Jwt

	expiresAt := time.Now().Add(time.Minute * time.Duration(jwtConfig.ValidWithinMinutes))
	uid, _ := uuid.NewV7()
	claims := JwtClaims{
		jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(id)),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uid.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtConfig.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
