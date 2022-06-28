package patients

import (
	"net/http"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) NewRecord(c *gin.Context) {
	var history models.MedicalHx
	var patient models.Patient
	session := sessions.Default(c)
	user_id := session.Get("user_id").(string)
	id := c.Param("id")

	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if result := h.DB.Preload("History").First(&patient, "id=?", id); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": result.Error,
		})
		return
	}
	history.User_id = user_id
	patient.History = append(patient.History, history)

	h.DB.Save(&patient)

	h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&patient)

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &patient,
	})

}
