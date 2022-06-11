package patients

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type handler struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func RegRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	h := &handler{
		DB:  db,
		RDB: rdb,
	}

	route := r.Group("/patients")
	route.GET("", h.GetPatients)
	route.GET("/:id", h.GetPatientById)
	route.POST("", h.NewPatientHandler)
	route.PUT("/:id", h.UpdatePatientHandler)
	route.DELETE("/:id", h.DeletePatient)

}
