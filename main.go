package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// TODO: Crear tabla
// [] El comando/datos

func main() {

	// NOTE: Abrir/crear la base de datos
	// verificar que funcione
	db, err := sql.Open("sqlite", "/animes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
