package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func refreshroute() {
	r := mux.NewRouter()

	r.HandleFunc("/refresh", models.Refresh).Methods("POST")

}
