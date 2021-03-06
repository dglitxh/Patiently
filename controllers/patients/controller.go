package patients

import (
	"fmt"

	"github.com/dglitxh/patiently/controllers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type handler struct {
	DB       *gorm.DB
	RDB      *redis.Client
	ClearRdb func()
}

func RegRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {

	DelRdb := func() {
		iter := rdb.Scan(rdb.Context(), 0, "", 0).Iterator()

		for iter.Next(rdb.Context()) {
			key := iter.Val()

			d, err := rdb.TTL(rdb.Context(), key).Result()
			if err != nil {
				panic(err)
			}

			if d == -1 { // -1 means no TTL
				if err := rdb.Del(rdb.Context(), key).Err(); err != nil {
					panic(err)
				}
			}
		}

		if err := iter.Err(); err != nil {
			panic(err)
		}
		fmt.Println("Redis cache cleared!")
	}

	h := &handler{
		DB:       db,
		RDB:      rdb,
		ClearRdb: DelRdb,
	}

	route := r.Group("/patients")
	route.GET("", h.GetPatients)
	route.GET("/:id", h.GetPatientById)

	auth := route.Use(middleware.AuthMiddleware())
	auth.POST("", h.NewPatientHandler)
	auth.POST("/:id/new_record", h.NewRecord)
	auth.PUT("/:id", h.UpdatePatientHandler)
	auth.GET("/:id/records", h.GetPtRecords)
	auth.DELETE("/:id", h.DeletePatient)

}
