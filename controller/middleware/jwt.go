package middleware

import (
	"os"

	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func JwtMiddleware() echo.MiddlewareFunc {
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	})
	return jwtMiddleware
}
