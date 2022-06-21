package middleware

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		claims := &models.Claims{}
		tokenVal := session.Get("token")

		if tokenVal == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "user not authenticated",
			})
			c.Abort()
		}

		tkn, err := jwt.ParseWithClaims(tokenVal.(string), claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(viper.Get("JWT_SECRET").(string)), nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"status": "Invalid token",
			})
			c.Abort()
		}

		if tkn == nil || !tkn.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"status": "Token is invalid.",
			})
			c.Abort()
		}

		c.Next()
	}
}
