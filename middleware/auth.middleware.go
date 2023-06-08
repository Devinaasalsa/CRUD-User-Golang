package middleware

import (
	"myapp-me/config"

	echojwt "github.com/labstack/echo-jwt/v4"
)

var IsAuthenticated = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte(config.SignatureKey),
})
