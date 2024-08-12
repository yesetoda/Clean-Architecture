package domain

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


func GenerateToken(username,role string)(string,error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return tokenString,err
}