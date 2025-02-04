package helpers

import (
	"fmt"
	"time"

	"github.com/akshay-kgen/todo-app/config"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(userId, email string) (string, error) {

	claims := &jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	}

	fmt.Println("claim", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.GetInstance().JwtSecret
	return token.SignedString([]byte(secretKey))
}
