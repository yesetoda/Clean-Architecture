package middleware

import "github.com/dgrijalva/jwt-go"

type GeneralAuth interface{
	Auth(tokenstring, secretkey string) jwt.Claims 
	AdminAuth(anyclaim any) bool 	
	UserAuth(anyclain any) bool 
}