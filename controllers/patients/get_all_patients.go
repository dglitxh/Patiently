package patients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dglitxh/patiently/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (h handler) GetPatients(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	page_size := c.DefaultQuery("page_size", "20")
	var patients []models.Patient

	cache, err := h.RDB.Get(h.RDB.Context(), "patients"+page).Result()

	if err != nil {
		fmt.Println("Request To DB")
		if results := h.DB.Scopes(Paginate(c.Request)).Find(&patients); results.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status": results.Error,
			})
			return
		}

		data, _ := json.Marshal(patients)
		h.RDB.Set(h.RDB.Context(), "patients"+page, string(data), 0)
		c.IndentedJSON(http.StatusOK, gin.H{
			"page":      page,
			"page size": page_size,
			"results":   &patients,
		})
	} else {
		fmt.Println("Request To Redis")
		json.Unmarshal([]byte(cache), &patients)
		c.IndentedJSON(http.StatusOK, gin.H{
			"page":      page,
			"page size": page_size,
			"results":   &patients,
		})
	}

}
