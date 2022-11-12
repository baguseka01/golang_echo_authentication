package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{"message": "Welcome to my website"})
}
