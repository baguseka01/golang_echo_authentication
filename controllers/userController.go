package controllers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/baguseka01/golang_echo_authentication/config"
	"github.com/baguseka01/golang_echo_authentication/middlewares"
	"github.com/baguseka01/golang_echo_authentication/models"
	"github.com/labstack/echo/v4"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return re.MatchString(email)
}

func Register(c echo.Context) error {
	var data map[string]interface{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if len(data["password"].(string)) <= 6 {
		return c.JSON(http.StatusInternalServerError, "Password cannot be smaller than 6 characters")
	}

	if !ValidateEmail(strings.TrimSpace(data["email"].(string))) {
		return c.JSON(http.StatusInternalServerError, "Email must be a-z 0-9")
	}

	var user models.User
	config.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&user)
	if user.ID != 0 {
		return c.JSON(http.StatusInternalServerError, "email already used")
	}

	userModel := models.User{
		Username: data["username"].(string),
		Email:    strings.TrimSpace(data["email"].(string)),
	}

	userModel.HashPassword(data["password"].(string))
	err := config.DB.Create(&userModel)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":    userModel,
		"message": "Congratulations, your registration was successful",
	})

}

func Login(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var user models.User
	config.DB.Where("email=?", data["email"]).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "your email is not registered")
	}

	if err := user.CheckPassword(data["password"]); err != nil {
		return c.JSON(http.StatusInternalServerError, "your password is wrong")
	}

	token, err := middlewares.GenerateJwtToken(user.Email, user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Generate JWT error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login successful",
		"user":    user,
		"token":   token,
	})
}
