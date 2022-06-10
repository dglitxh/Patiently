package patients

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if result := h.DB.First(&patient, id); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": result.Error,
		})
		return
	}

	h.DB.Delete(&patient)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "Patient with id " + id + " has been deleted",
	})
}
