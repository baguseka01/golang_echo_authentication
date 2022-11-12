package auth

import (
	"net/http"
	"strings"

	"github.com/baguseka01/golang_echo_authentication/middlewares"
	"github.com/labstack/echo/v4"
)

func Auth(context echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		authHeader := e.Request().Header.Get("Authorization")
		bearerToken := (strings.Split(authHeader, " "))[1]
		if bearerToken == "" {
			return e.JSON(http.StatusUnauthorized, &echo.Map{
				"message": "Unauthorized",
			})
		}

		err := middlewares.ValidateToken(bearerToken)
		if err != nil {
			return e.JSON(http.StatusUnauthorized, &echo.Map{
				"message": err.Error(),
			})
		}

		return context(e)
	}
}
