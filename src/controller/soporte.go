package controller

import (
	"context"
	"fmt"
	"net/http"
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No se pudo procesar el parametro de la URL."})
		return
	}

	// Obtener el token para notificar al usuario
	consulta, result := service.ObtenerTokenEstudiante(idConsulta)
	if result != "" && result != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result})
		return
	}

	// Validar si el alumno tiene un token
	if consulta.TmUsuario.TokenApp == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "El alumno no tiene un token, comuníquese con el área de informática."})
		return
	}

	// Configurar Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile("src/helper/keys/app-push-notification-f632c-firebase-adminsdk-d271n-f5a3ed91dc.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Se genero un problema al envíar la notificación"})
		return
	}

	// Crear un cliente de FCM
	client, err := app.Messaging(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Se genero un problema al envíar la notificación"})
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
		// Token: "fJUEIvR0njdVJgshHOSZ_S:APA91bEu4l2Aip7iKs-cCu0rW5bqiaveks2RlUxWkfZTWPsZma8ZW4qss73mNxsB_g9dtQh5L35MSqoVLAwhl9lXUJziRAbpRfgSssucuYgq8ywmeGx1KQ5KjfEfqvtSyN2ohlvK-ep_",
		Token: consulta.TmUsuario.TokenApp,
	}

	// Envía el mensaje de notificación push
	response, err := client.Send(ctx, message)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Se genero un problema al envíar la notificación"})
		return
	}

	fmt.Println(response)

	c.IndentedJSON(http.StatusOK, "Envió correctamente la notificación.")
}
