package service

import (
	"context"
	"database/sql"
	"servicio-pushnotificacion/src/database"
	"servicio-pushnotificacion/src/model"
)

var contx = context.Background()

func ObtenerTokenEstudiante(idConsulta string) (model.Consulta, string) {
	consulta := model.Consulta{}

	db, err := database.CreateConnection()
	if err != nil {
		return consulta, "No se puedo establecer la conexión al servidor."
	}

	defer db.Close()

	query := `SELECT 
	co.ticket,
	co.asunto, 
	co.descripcion,
	tu.tokenApp 
	FROM Soporte.Consulta  AS co 
	INNER JOIN Est_Estudiante AS es ON es.Est_Id = co.Est_Id
	INNER JOIN seguridad.TM_Usuario AS tu ON tu.c_cod_usuario = es.Est_Id
	WHERE co.idConsulta = @idConsulta`

	row := db.QueryRowContext(contx, query, sql.Named("idConsulta", idConsulta))

	err = row.Scan(&consulta.Ticket, &consulta.Asunto, &consulta.Descripcion, &consulta.TmUsuario.TokenApp)
	if err != nil {
		return consulta, "No se puedo leer la fila requerida."
	}

	return consulta, "ok"
}
