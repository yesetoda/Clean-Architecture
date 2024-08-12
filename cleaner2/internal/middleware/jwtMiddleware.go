package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth struct {
}

func NewJWTAuth() *JWTAuth {
	return &JWTAuth{}
}
func (j *JWTAuth) Auth(tokenstring, secretkey string) jwt.Claims {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, nil
		}
		return []byte(secretkey), nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	// Set the token claims to the context
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

func (j *JWTAuth) AdminAuth(anyclaims any) bool {
	claims := anyclaims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role == "admin"
}

func (j *JWTAuth) UserAuth(anyclaims any) bool {
	claims := anyclaims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role == "admin" || role == "user"
}
