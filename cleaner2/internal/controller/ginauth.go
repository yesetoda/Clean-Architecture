package controller

import (
	"example/cleaner2/internal/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (tc *GInGenaralController) AuthMiddlewareGIn(auth middleware.GeneralAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = os.Getenv("JWT_KEY")
		tokenString := c.GetHeader("Authorization")
		claims := auth.Auth(tokenString, secretKey)
		fmt.Println("this is the claim",claims)
		if claims != nil {
			c.Set("claims", claims)
			c.Next() // Proceed to the next handler if authorized
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort() // Stop further processing if unauthorized
	}

}

func (tc *GInGenaralController) AdminMiddlewareGin(auth middleware.GeneralAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// handle the panic
				// fmt.Println("Recovered from panic:", r)
				c.JSON(http.StatusForbidden, gin.H{"error": "UnAuthorized", "message": "must be an admin to do such task"})
				c.Abort()
			}
		}()
		claims := c.MustGet("claims")
		if !auth.AdminAuth(claims) {
			c.JSON(http.StatusForbidden, gin.H{"error": "UnAuthorized", "message": "must be an admin to do such task"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (tc *GInGenaralController) UserMiddlewareGin(auth middleware.GeneralAuth) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// handle the panic
				// fmt.Println("Recovered from panic:", r)
				c.JSON(http.StatusForbidden, gin.H{"error": "UnAuthorized", "message": "must be an admin to do such task"})
				c.Abort()
			}
		}()
		claims := c.MustGet("claims")
		if !auth.UserAuth(claims) {
			c.JSON(http.StatusForbidden, gin.H{"error": "you must log in first"})
			c.Abort()
			return
		}
		c.Next()
	}
}
