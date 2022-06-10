package patients

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	route := r.Group("/patients")
	route.GET("", h.GetPatients)
	route.GET("/:id", h.GetPatientById)
	route.POST("", h.NewPatientHandler)
	route.PUT("/:id", h.UpdatePatientHandler)
	route.DELETE("/:id", h.DeletePatient)

}
