package Routes

import (
	"restapi/models"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/books/create", models.CreateBookapi).Methods("POST")
	r.HandleFunc("/books", models.GetBooksapi).Methods("GET")
	r.HandleFunc("/books/{id}", models.GetBookapi).Methods("GET")
	r.HandleFunc("/books/{id}", models.UpdateBookapi).Methods("POST")
	r.HandleFunc("/books/{id}", models.DeleteBookapi).Methods("DELETE")
}
