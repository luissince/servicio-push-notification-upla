package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"servicio-pushnotificacion/src/helper"
	"servicio-pushnotificacion/src/service"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

func EnviarNotificacion(c *gin.Context) {
	// Obtener el idConsulta de la URL
	idConsulta := c.Param("idConsulta")
	if idConsulta == "" {
		log.Println("No se pudo procesar el parametro de la URL.")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No se pudo procesar el parametro de la URL."})
		return
	}

	// Obtener el token para notificar al usuario
	consulta, result := service.ObtenerTokenEstudiante(idConsulta)
	if result != "" && result != "ok" {
		log.Println(result)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result})
		return
	}

	// Validar si el alumno tiene un token
	if consulta.TmUsuario.TokenApp == "" {
		log.Println("El alumno no tiene un token, comuníquese con el área de informática.")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "El alumno no tiene un token, comuníquese con el área de informática."})
		return
	}

	fmt.Println(consulta.TmUsuario.TokenApp)

	var ruta_firebase string = os.Getenv("RUTA_FIREBASE")
	fmt.Println(ruta_firebase)

	// Configurar Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(ruta_firebase)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Se genero un problema al carga el archivo de configuración."})
		return
	}

	// Crear un cliente de FCM
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Se genero un problema al crear al cliente de mensaje FCM."})
		return
	}

	// Creamos un mapa donde agregar que tipo de pantalla se va abrir al hacerle click a la notificación
	data := make(map[string]string)
	data["idConsulta"] = idConsulta
	data["tipo"] = "centroayuda"

	// Crear el mensaje de notificación push
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Su consulta del Ticket: N° - " + helper.GenerateZeros(consulta.Ticket, 6) + " ha sido atendido, revise por favor.",
			Body:  consulta.Asunto,
		},
		Data: data,
		//Token: "fJUEIvR0njdVJgshHOSZ_S:APA91bHruGtQgzzPh9uNA1pUonpkh8crqtwCqpkvRR2_YEKHtVwS6u8pmuPWK3H7EfmH8NPyilDlj82E_92vzEkj0F1heW4q7ayD9ovu6TQe18xABe2ahCWMHKxAeDn835KArluaI2MZ",
		Token: consulta.TmUsuario.TokenApp,
	}

	// Envía el mensaje de notificación push
	response, err := client.Send(ctx, message)
	if err != nil {
		if messaging.IsRegistrationTokenNotRegistered(err) {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Token de registro no válido o no registrado."})
			return
		}

		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error al enviar la notificación."})
		return
	}

	fmt.Println(response)

	c.IndentedJSON(http.StatusOK, "Envió correctamente la notificación.")
}

func ValidarNotificacion(c *gin.Context) {
	idConsulta := c.Param("idConsulta")
	if idConsulta == "" {
		log.Println("No se pudo procesar el parametro de la URL.")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No se pudo procesar el parametro de la URL."})
		return
	}

	c.IndentedJSON(http.StatusOK, "Envió correctamente la notificación del id:"+idConsulta)
}
