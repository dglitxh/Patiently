package patients

import (
	"errors"
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) UpdateRecords(c *gin.Context) {
	var record models.Records
	var body models.MedicalHx
	var patient models.Patient
	id := c.Param("id")

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if result := h.DB.First(&patient, "id=?", id); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": result.Error,
		})
		return
	}

	result := h.DB.First(&record, "patient=?", &patient)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "Record already available for this patient.",
		})
		return
	}

	record.History = append(record.History, body)

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &record,
	})
}
