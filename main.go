package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Jwt-Project/controllers"
	"github.com/seheraksam/Jwt-Project/initializers"
)

func init() {
	if err := initializers.LoadEnvVariables(); err != nil {
		log.Fatal("Env yüklenemedi:", err)
	}
	if err := initializers.ConnectToDb(); err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}
	if err := initializers.SyncDatabase(); err != nil {
		log.Fatal("Database senkronizasyon hatası:", err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signUp", controllers.SignUp)
	r.Run()
}
