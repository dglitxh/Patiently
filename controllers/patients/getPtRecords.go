package patients

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPtRecords(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if result := h.DB.Preload("History").First(&patient, "id=?", id); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": result.Error,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &patient,
	})
}
