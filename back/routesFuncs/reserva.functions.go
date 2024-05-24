package routesfuncs

import (
	"encoding/json"
	"net/http"

	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/db"
	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/models"
)

func ReservationsHandler(w http.ResponseWriter, req *http.Request) {

	var reserva models.Reserva

	json.NewDecoder(req.Body).Decode(&reserva) //Decodificacion del json enviado desde el cliente al servidor

	//Guardado en DB

	createdReserva := db.DB.Create(&reserva) //Se trata de guardar la reserva en DB

	err := createdReserva.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //codigo 400
		w.Write([]byte(err.Error()))

		return
	}

	//luego, se debe cambiar el estado de dicho carro a TRUE

	db.DB.Model(&models.Carro{}).Where("estado_reserva = ? AND carro_id = ?", false, reserva.Carro_ID).Update("estado_reserva", true)

	json.NewEncoder(w).Encode(&reserva) //Se envia la reserva

}

func DeleteReservationsHandler(w http.ResponseWriter, req *http.Request) {

	var reserva models.Reserva

	var findReserva models.Reserva

	json.NewDecoder(req.Body).Decode(&reserva)

	reservaID := reserva.ReservaID

	db.DB.Where("reserva_id = ?", reservaID).First(&findReserva)

	if findReserva.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Reserva NO encontrada"))

		return

	}

	db.DB.Model(&models.Carro{}).Where("carro_id = ?", findReserva.Carro_ID).Update("estado_reserva", false)

	db.DB.Unscoped().Delete(&findReserva) //Eliminación de la reserva
	w.WriteHeader(http.StatusNoContent)   //204

}

func ReportHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Reporte"))

}
