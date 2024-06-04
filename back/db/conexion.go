package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConexionDB() {

	var err error

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	var DSN = "host=db user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil { //Si existe algun error

		log.Fatal(err)

	} else {

		log.Println("Conexion a DB exitosa")
	}

}
