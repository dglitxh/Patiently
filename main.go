package main

import (
	"net/http"

	"github.com/dglitxh/patiently/common/db"
	"github.com/dglitxh/patiently/controllers/patients"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)



func main() {
	viper.SetConfigFile("./common/envs/.env")
	viper.ReadInConfig()

	db_url := viper.Get("DB_URL").(string)
	port := viper.GetString("PORT")

	r := gin.Default()
	h := db.InitDb(db_url)

	patients.RegRoutes(r, h)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "We live bruh",
		})
	})

	r.Run(port)
}
