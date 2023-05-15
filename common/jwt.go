package common

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/exception"
)

func GenerateToken(email, noTelp string, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"email":   email,
		"no_telp": noTelp,
		"exp":     time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
