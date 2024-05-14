package routesfuncs

import (
	"net/http"
)

func ReservationsHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Reservas"))

}

func ReportHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Reporte"))

}
