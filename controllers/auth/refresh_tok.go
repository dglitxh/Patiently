package auth

import (
	"net/http"
	"time"

	"github.com/dglitxh/patiently/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RefreshJwt(c *gin.Context) {
	claims := &models.Claims{}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token is not expired yet",
		})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.Get("JWT_SECRET").(string)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	session := sessions.Default(c)
	session.Set("token", tokenString)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "token updated succesfully.",
	})

}
