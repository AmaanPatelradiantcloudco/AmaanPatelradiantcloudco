package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func homeroute() {
	r := mux.NewRouter()

	r.HandleFunc("/home", models.Homeapi).Methods("GET")

}
