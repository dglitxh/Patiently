package patients

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
)

func (h handler) NewPatientHandler(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"server": "not working",
		})
		return
	}

	if result := h.DB.Create(&patient); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": result.Error,
		})
		return
	}

	h.ClearRdb()

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &patient})
}
