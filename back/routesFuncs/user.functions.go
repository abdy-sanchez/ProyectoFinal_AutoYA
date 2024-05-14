package routesfuncs

import (
	"encoding/json"
	"net/http"

	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/db"
	"github.com/abdy-sanchez/ProyectoFinal_AutoYA/models"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {

	var user models.Usuario

	var findUser models.Usuario

	json.NewDecoder(req.Body).Decode(&user)

	email := user.Email //Contraseña e Email ingresados por el cliente

	passw := user.Password

	db.DB.Where("email = ? AND password = ?", email, passw).First(&findUser)

	if findUser.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Usuario no encontrado"))

		return
	}

	json.NewEncoder(w).Encode(&findUser)

}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {

	var user models.Usuario

	json.NewDecoder(req.Body).Decode(&user) //Decodificacion del json enviado desde el cliente al servidor

	//Guardado en DB

	createdUser := db.DB.Create(&user)

	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //codigo 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user) //Se envia el usuario creado

}

func UsersHandler(w http.ResponseWriter, req *http.Request) {

	var users []models.Usuario

	db.DB.Find(&users) //Buscamos toda la tabla Usuarios

	json.NewEncoder(w).Encode(&users) //Retorna todas las filas de la tabla

}
