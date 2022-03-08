package handlers

import (
	"log"              //provide logging features (debug,error,warning)
	"net/http"         //used to get GET request specifically
	"restapi/Database" //importing database over here since it is using function initialmigration which is there in database directory

	"github.com/gorilla/handlers" //this package is for handling purpose which is for methods and cors functionality
	"github.com/gorilla/mux"      //used for matching request and to their respective handler
)

func initializeRouter() {
	r := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})                                        //cors functionality the request is made from
	originsOk := handlers.AllowedOrigins([]string{"*"})                                                       //this is taken from cors.go
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "POST", "OPTIONS"}) //we are giving different methods that are required
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
func main_api() { //we are declaring main function
	Database.InitialMigration() //taken from database initialmigration which is taken from package gorm
	initializeRouter()
}
