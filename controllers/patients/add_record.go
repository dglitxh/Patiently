package patients

import (
	"net/http"

	// "errors"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h handler) NewRecord(c *gin.Context) {
	var record models.Records
	var history models.MedicalHx
	var user models.User
	session := sessions.Default(c)
	user_id := session.Get("user_id").(string)
	id := c.Param("id")

	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	var patient models.Patient
	if result := h.DB.First(&patient, "id=?", id); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": result.Error,
		})
		return
	}

	// result := h.DB.First(&record, "patient=?", id)
	// if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"status": "Record already available for this patient.",
	// 	})
	// 	return
	// }

	if res := h.DB.First(&user, "id=?", user_id); res.Error != nil {
		c.JSON(http.StatusBadRequest, res.Error)
	}
	record.History = append(record.History, history)
	record.Patient = patient
	record.User = user

	if result := h.DB.Create(&record); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "could not create record",
		})
		return
	}

	h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&record)

	c.IndentedJSON(http.StatusOK, gin.H{
		"result": &record,
	})

}
