package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func booksroute() {
	r := mux.NewRouter()
	r.HandleFunc("/books/create", models.CreateNewBookapi).Methods("POST")
	r.HandleFunc("/books", models.GetAllBooksapi).Methods("GET")
	r.HandleFunc("/books/{id}", models.GetSingleBookapi).Methods("GET")
	r.HandleFunc("/books/{id}", models.UpdateExistingBookapi).Methods("POST")
	r.HandleFunc("/books/{id}", models.DeletePreviousBookapi).Methods("DELETE")
}
