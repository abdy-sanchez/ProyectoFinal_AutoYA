package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost user=abdy password=password dbname=autoya port=5432"

var DB *gorm.DB

func ConexionDB() {

	var err error

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil { //Si existe algun error

		log.Fatal(err)

	} else {

		log.Println("Conexion a DB exitosa")
	}

}
