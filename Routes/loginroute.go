package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func loginroute() {
	r := mux.NewRouter()

	r.HandleFunc("/login", models.Loginapi).Methods("POST")

}
