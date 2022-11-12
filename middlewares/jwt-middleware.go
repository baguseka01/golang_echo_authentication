package middlewares

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJwtToken(email string, username string) (string, error) {
	if os.Getenv("JWT_SECRET_KEY") == "" {
		log.Fatal("Error loading JWT_SECRET_KEY .env file")
	}

	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	expirationTime := time.Now().Add(10 * time.Minute)
	Claims := &Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(signedToken string) error {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return errors.New("invalid token")
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
