package main

import (
	//"encoding/json"
	"log"
	"net/http"

	routesfuncs "github.com/abdy-sanchez/ProyectoFinal_AutoYA/routesFuncs"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter() //Declaración del enrutador

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
	log.Fatal(http.ListenAndServe(":8080", router))

}
