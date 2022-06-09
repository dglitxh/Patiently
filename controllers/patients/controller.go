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
	route.POST("/", h.NewPatientHandler)
}
