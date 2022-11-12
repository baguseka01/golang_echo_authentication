package routes

import (
	"github.com/baguseka01/golang_echo_authentication/auth"
	"github.com/baguseka01/golang_echo_authentication/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure middleware with the custom claims type
	jwt := e.Group("/jwt", auth.Auth)
	jwt.GET("/home", controllers.Home)

	return e
}
