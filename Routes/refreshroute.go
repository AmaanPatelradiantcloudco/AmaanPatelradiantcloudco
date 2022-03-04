package Routes

import (
	"restapi/Database"
	"restapi/models"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/refresh", models.Refresh).Methods("POST")

}
func main_api() {
	Database.InitialMigration()
	initializeRouter()
}
