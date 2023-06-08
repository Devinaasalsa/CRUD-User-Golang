package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetCORS() echo.MiddlewareFunc {
	config := middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3001",
		},
		AllowMethods: []string{
			echo.GET,
			echo.PUT,
			echo.POST,
			echo.DELETE,
			echo.PATCH,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderCookie,
			echo.HeaderAccept,
		},
	}

	return middleware.CORSWithConfig(config)
}
