package routesfuncs

import (
	"encoding/json"
	"net/http"

	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/db"
	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/models"
	"github.com/gorilla/mux"
)

func GetCarsHandler(w http.ResponseWriter, req *http.Request) {

	var cars []models.Carro

	db.DB.Find(&cars) //Buscamos toda la tabla Carros

	json.NewEncoder(w).Encode(&cars) //Retorna todas las filas y columnas de la tabla

}

func GetCarsHandlerFiltrado(w http.ResponseWriter, req *http.Request) {

	var findCar_filtrado []models.Carro

	parametros := mux.Vars(req)

	db.DB.Where(&models.Carro{Marca: parametros["marca"], Tipo: parametros["tipo"], Combustible: parametros["combustible"], Transmision: parametros["transmision"]}).Find(&findCar_filtrado)

	json.NewEncoder(w).Encode(&findCar_filtrado)

}
