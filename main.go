package main

import (
	"fmt"
	"net/http"

	"github.com/dglitxh/patiently/common/db"
	"github.com/dglitxh/patiently/controllers/auth"
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
	redi := db.InitRdb()

	fmt.Println(redi.Ping(redi.Context()))
	patients.RegRoutes(r, h, redi)
	auth.RegAuthRoutes(r, h)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "We are LIVE!!!"})
	})

	r.Run(port)
}
