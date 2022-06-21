package auth

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-contrib/sessions"
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
			"status": "Authentication failed: wrong email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"status": "Authentication failed: wrong email or password",
		})
		return
	}

	secret := viper.GetString("JWT_Secret")

	JWT := SignJWT(user.Username, secret)

	session := sessions.Default(c)
	session.Set("token", JWT.Token)
	session.Set("email", user.Email)
	session.Set("username", user.Username)
	session.Save()

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "user authenticated succesfully",
	})
}
