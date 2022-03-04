package Routes

import (
	"restapi/Database"
	"restapi/models"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/register", models.CreateUserapi).Methods("POST")

}
func main_api() {
	Database.InitialMigration()
	initializeRouter()
}
