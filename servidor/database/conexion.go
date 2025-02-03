package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Conectar() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/marketplace-base?parseTime=true"

	var err error
	// ¡Usa asignación a la variable global, no := !
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}

	// ¡Quita el defer DB.Close()! La conexión debe permanecer abierta.

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	fmt.Println("Conexión exitosa")
}
