package middleware

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenVal := c.GetHeader("Authorization")
		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenVal, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(viper.Get("JWT_SECRET").(string)), nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"status": err.Error(),
			})
			panic(err.Error())
		}

		if tkn == nil || !tkn.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"status": "Token is invalid.",
			})
			panic("Invalid token")
		}

		c.Next()
	}
}
