package main

import (
	"log"
	"net/http"
	"restapi/docs"
	"restapi/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
// @title Books API
// @version 1.0
// @description This is a Service for Managing Books
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath /
func initializeRouter() {
	r := mux.NewRouter()
	//for JWT
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	//  r.PathPrefix("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:1323/swagger/doc.json").Methods("GET")
	r.HandleFunc("/login", models.Loginapi).Methods("POST")
	r.HandleFunc("/home", models.Homeapi).Methods("GET")
	r.HandleFunc("/refresh", models.Refresh).Methods("POST")
	// r.HandleFunc("/logout", models.Logout).Methods("POST")
	//For USERS
	r.HandleFunc("/users", models.GetUser).Methods("GET")
	//r.HandleFunc("/users/{id}", models.GetUser).Methods("GET")
	// r.HandleFunc("/users/{id}", models.DeleteUser).Methods("DELETE")
	// r.HandleFunc("/users/{id}", models.UpdateUser).Methods("PATCH")
	r.HandleFunc("/register", models.CreateUserapi).Methods("POST")
	//  //For Authors
	r.HandleFunc("/authors", models.GetAuthorsapi).Methods("GET")
	r.HandleFunc("/authors/{id}", models.GetAuthorapi).Methods("GET")
	r.HandleFunc("/authors/{id}", models.DeleteAuthorapi).Methods("DELETE")
	r.HandleFunc("/authors/{id}", models.UpdateAuthorapi).Methods("PATCH")
	r.HandleFunc("/authors/create", models.CreateAuthorapi).Methods("POST")
	//  //For Books
	r.HandleFunc("/books/create", models.CreateBookapi).Methods("POST")
	r.HandleFunc("/books", models.GetBooksapi).Methods("GET")
	r.HandleFunc("/books/{id}", models.GetBookapi).Methods("GET")
	r.HandleFunc("/books/{id}", models.UpdateBookapi).Methods("PATCH")
	r.HandleFunc("/books/{id}", models.DeleteBookapi).Methods("DELETE")
	docs.SwaggerInfoabc_swagger.Host = "localhost:8080"
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "POST", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
	//log.Fatal(http.ListenAndServe(":8080",r))
	// }
}
func main() {
	models.InitialMigration()
	initializeRouter()
}

// \ //
