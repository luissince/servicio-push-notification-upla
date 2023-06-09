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

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 		c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		c.Next()
// 	}
// }

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
	//router.Use(corsMiddleware())
	router.Use(cors.Default())
	// config.AllowOrigins = []string{"*"}
	// router.Use(cors.New(config))

	// Rutas
	router.GET("/", func(c *gin.Context) {
		log.Println("Endpoint ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "API GO LANG",
		})
	})

	router.GET("/notificar/:idConsulta", controller.EnviarNotificacion)

	router.Run(go_port)
}
