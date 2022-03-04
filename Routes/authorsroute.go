package Routes

import (
	"restapi/Database"
	"restapi/models"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	//  //For Authors
	r.HandleFunc("/authors", models.GetAuthorsapi).Methods("GET")
	r.HandleFunc("/authors/{id}", models.GetAuthorapi).Methods("GET")
	r.HandleFunc("/authors/{id}", models.DeleteAuthorapi).Methods("DELETE")
	r.HandleFunc("/authors/{id}", models.UpdateAuthorapi).Methods("POST")
	r.HandleFunc("/authors/create", models.CreateAuthorapi).Methods("POST")

}
func main_api() {
	Database.InitialMigration()
	initializeRouter()
}
