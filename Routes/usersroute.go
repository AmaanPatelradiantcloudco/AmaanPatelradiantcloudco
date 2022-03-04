package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", models.GetUser).Methods("GET")

}
