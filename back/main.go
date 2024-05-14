package main

import (
	//"encoding/json"
	"log"
	"net/http"

	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/db"
	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/models"
	routesfuncs "github.com/abdy-sanchez/ProyectoFinal_AutoYA/routesFuncs"
	"github.com/gorilla/mux"
)

func main() {

	//Conexion a base de datos

	db.ConexionDB()

	//Migraciones de las tablas para la DB

	db.DB.AutoMigrate(models.Usuario{})
	db.DB.AutoMigrate(models.Carro{})
	db.DB.AutoMigrate(models.Reserva{})

	//Declaración del enrutador

	router := mux.NewRouter()

	//Funcion para mostrar todos los usuarios registrados

	router.HandleFunc("/users", routesfuncs.UsersHandler).Methods("GET")

	//Funcion para Logeo
	router.HandleFunc("/login", routesfuncs.LoginHandler).Methods("POST")

	//Funcion para Registro
	router.HandleFunc("/register", routesfuncs.RegisterHandler).Methods("POST")

	//Funcion para Buscar un carro en la DB
	router.HandleFunc("/cars", routesfuncs.GetCarsHandler).Methods("GET")

	//Funcion para Reservar un carro
	router.HandleFunc("/reservations", routesfuncs.ReservationsHandler).Methods("POST")

	//Funcion para ver el reporte de reservas del usuario
	router.HandleFunc("/report", routesfuncs.ReportHandler).Methods("POST")

	//Inicializacion del Servidor
	log.Println("servidor en puerto :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
