package auth

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if results := h.DB.Find(&user, "id=?", id); results.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": results.Error,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, &user)
}
