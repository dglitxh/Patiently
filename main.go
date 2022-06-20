package main

import (
	"fmt"
	"net/http"

	"github.com/dglitxh/patiently/common/db"
	"github.com/dglitxh/patiently/controllers/auth"
	"github.com/dglitxh/patiently/controllers/patients"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./common/envs/.env")
	viper.ReadInConfig()

	db_url := viper.Get("DB_URL").(string)
	port := viper.GetString("PORT")
	secret := viper.GetString("STORE_SECRET")

	r := gin.Default()
	h := db.InitDb(db_url)
	redi := db.InitRdb()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(secret))
	r.Use(sessions.Sessions("patiently", store))
	patients.RegRoutes(r, h, redi)
	auth.RegAuthRoutes(r, h)

	fmt.Println(redi.Ping(redi.Context()))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "We are LIVE!!!"})
	})

	r.Run(port)
}
