package routesfuncs

import (
	"net/http"
)

func GetCarsHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("DB Autos"))

}
