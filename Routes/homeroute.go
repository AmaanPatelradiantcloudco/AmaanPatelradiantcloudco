package Routes

import (
	"restapi/Database"
	"restapi/models"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/home", models.Homeapi).Methods("GET")

}
func main_api() {
	Database.InitialMigration()
	initializeRouter()
}
