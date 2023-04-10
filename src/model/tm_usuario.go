package model

type TmUsuario struct {
	C_cod_usuario string `json:"c_cod_usuario"`
	C_pas_usuario string `json:"c_pas_usuario"`
	TokenApp      string `json:"tokenApp"`
}
