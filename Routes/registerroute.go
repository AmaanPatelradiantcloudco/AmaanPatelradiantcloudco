package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func registerroute() {
	r := mux.NewRouter()

	r.HandleFunc("/register", models.CreateNewUserapi).Methods("POST")

}
