package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"servicio-pushnotificacion/src/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	time.LoadLocation("America/Lima")
	godotenv.Load()

	var go_port string = os.Getenv("GO_PORT")

	var ruta_log string = os.Getenv("RUTA_LOG")

	// Crear archivo log
	f, err := os.Create(ruta_log)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	router := gin.Default()
	router.Use(gin.Logger())

	// Middleware para CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// Rutas
	router.GET("/", func(c *gin.Context) {
		log.Println("Endpoint ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG",
		})
	})

	router.GET("/consulta/:idConsulta", controller.EnviarNotificacion)
	// router.GET("/user/:idUsuario", service.GetUserById)
	// router.POST("/user", service.InsertUser)
	// router.PUT("/user", service.UpdateUser)
	// router.DELETE("/user", service.DeleteUsuario)

	router.Run(go_port)
}
