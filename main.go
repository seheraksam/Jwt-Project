package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Jwt-Project/controllers"
	"github.com/seheraksam/Jwt-Project/initializers"
	"github.com/seheraksam/Jwt-Project/middleware"
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
	r.Use(middleware.RequireAuth)
	r.POST("/signUp", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/validate", controllers.Validate)
	r.Run()
}
