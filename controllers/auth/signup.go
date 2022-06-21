package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-contrib/sessions"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) SignupHandler(c *gin.Context) {
	var creds models.User

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if len(creds.Email) < 1 || len(creds.Password) < 1 || len(creds.Username) < 1 {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"status": "Credentials not provided",
		})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	creds.Password = string(password)
	email := creds.Email

	result := h.DB.First(&creds, "email=?", email)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "Email is already being used.",
		})
		return
	}

	creds.Id = xid.New().String()
	creds.CreatedAt = time.Now()

	if result := h.DB.Create(&creds); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": result.Error,
		})
		return
	}

	secret := viper.Get("JWT_Secret").(string)

	JWT := SignJWT(creds.Username, secret)

	session := sessions.Default(c)
	session.Set("token", JWT.Token)
	session.Set("email", creds.Email)
	session.Set("username", creds.Username)
	session.Save()

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "new user authenticated succesfully",
	})
}
