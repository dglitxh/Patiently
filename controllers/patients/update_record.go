package patients

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateRecords(c *gin.Context) {
	var record models.Records
	// var body models.MedicalHx
	id := c.Param("id")

	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status": err.Error(),
	// 	})
	// 	return
	// }

	err := h.DB.Model(&record).Preload("History").Find(&record, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &record,
	})
}
