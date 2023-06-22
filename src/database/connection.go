package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

// connectionString construye y devuelve la cadena de conexión a la base de datos.
func connectionString() string {
	// Carga las variables de entorno desde el archivo .env.
	godotenv.Load()

	// Obtiene los valores de las variables de entorno para la configuración de la base de datos.
	server := os.Getenv("SERVER_DB")
	port := os.Getenv("PORT_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	database := os.Getenv("NAME_DB")

	// Construye la cadena de conexión usando los valores obtenidos de las variables de entorno.
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", server, user, password, port, database)

	return connString // Retorna la cadena de conexión construida.
}

// CreateConnection crea y devuelve una conexión a la base de datos.
func CreateConnection() (*sql.DB, error) {
	// Abre una conexión a la base de datos utilizando la cadena de conexión.
	db, err := sql.Open("sqlserver", connectionString())
	if err != nil {
		// Si ocurre un error al abrir la conexión, retorna el error.
		return nil, err
	}

	// Establece el número máximo de conexiones abiertas permitidas.
	db.SetMaxOpenConns(5)

	// Realiza un ping a la base de datos para comprobar la conexión.
	err = db.Ping()
	if err != nil {
		// Si ocurre un error al realizar el ping, retorna el error.
		return nil, err
	}

	return db, nil // Retorna la conexión a la base de datos sin errores.
}
