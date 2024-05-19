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

	var car_filtrado []models.Carro

	parametros := mux.Vars(req)

	db.DB.Where("estado_reserva = ? AND marca = ?", false, parametros["marca"]).Find(&car_filtrado)

	json.NewEncoder(w).Encode(&car_filtrado)

}
