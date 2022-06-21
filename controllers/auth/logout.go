package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "user logged out succesfully",
	})
}
