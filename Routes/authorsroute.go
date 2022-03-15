package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func authorsroute() {
	r := mux.NewRouter()

	//  //For Authors
	r.HandleFunc("/authors", models.GetAllAuthorsapi).Methods("GET")           //for getting all the authors
	r.HandleFunc("/authors/{id}", models.GetSingleAuthorapi).Methods("GET")    //getting author with id
	r.HandleFunc("/authors/{id}", models.DeleteAuthorapi).Methods("DELETE")    //deleting the existing authors
	r.HandleFunc("/authors/{id}", models.UpdateAuthorapi).Methods("PATCH")     //update the author which already existed
	r.HandleFunc("/authors/create", models.CreateNewAuthorapi).Methods("POST") ///for creating new author using the method post

}
