package routesfuncs

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("logeo"))

}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Registro"))

}

func GetCarsHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("DB Autos"))

}

func ReservationsHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Reservas"))

}

func ReportHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Reporte"))

}
