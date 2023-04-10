package main

import (
	// "fmt"
	"servicio-pushnotificacion/src/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middleware para CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/consulta/:idConsulta", controller.EnviarNotificacion)
	// router.GET("/user/:idUsuario", service.GetUserById)
	// router.POST("/user", service.InsertUser)
	// router.PUT("/user", service.UpdateUser)
	// router.DELETE("/user", service.DeleteUsuario)

	router.Run("localhost:8080")
}
