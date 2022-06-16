package auth

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) LoginHandler(c *gin.Context) {
	var user models.User
	var creds models.Login

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if res := h.DB.First(&user, "email=?", &creds.Email); res.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": "Authentication failed",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"status": "crypt: Authentication failed",
		})
		return
	}

	secret := viper.Get("JWT_Secret").(string)

	JWT := SignJWT(user.Username, secret)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":       &user.Id,
		"username": &user.Username,
		"JWT":      JWT,
	})
}
