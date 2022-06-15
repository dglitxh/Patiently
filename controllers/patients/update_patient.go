package patients

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdatePatientHandler(c *gin.Context) {
	id := c.Param("id")
	var body models.Patient

	if err := c.ShouldBindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	var patient models.Patient
	if result := h.DB.First(&patient, id); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": result.Error,
		})
		return
	}

	patient.Name = body.Name
	patient.DOB = body.DOB
	patient.Insurance = body.Insurance
	patient.Gender = body.Gender
	patient.Occupation = body.Occupation

	h.DB.Save(&patient)
	h.ClearRdb()

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &patient,
	})
}
