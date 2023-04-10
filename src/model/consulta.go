package model

type Consulta struct {
	Ticket        int    `json:"ticket"`
	Asunto        string `json:"asunto"`
	TipoConsulta  int    `json:"tipoConsulta"`
	Descripcion   string `json:"descripcion"`
	Estado        int    `json:"estado"`
	Fecha         string `json:"fecha"`
	Hora          string `json:"hora"`
	Est_Id        string `json:"est_id"`
	C_cod_usuario string `json:"cod_usuario"`

	EstEstudiante EstEstudiante
	TmUsuario     TmUsuario
}
