package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
type _Book struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CreateBookRequest struct {
	Name        string `json:"book_name"`
	Description string `json:"descript"`
	Writtenat   string `json:"writtenat"`
}

// type CreateBookResponse struct {
//  Name        string `json:"book_name"`
//  Description string `json:"descript"`
// }
// type GetBooksRequest struct {
//  BookID string `json:"book_id"`
// }
// type GetBooksResponse struct {
//  BookID string `json:"book_id"`
// }
// type GetBookRequest struct {
//  Name        string `json:"book_name"`
//  Description string `json:"descript"`
// }
// type GetBookResponse struct {
//  BookID string `json:"book_id"`
// }
// type UpdateBookRequest struct {
//  Name        string `json:"book_name"`
//  Description string `json:"descript"`
// }
// type UpdateBookResponse struct {
//  BookID string `json:"book_id"`
// }
// type DeleteBookRequest struct {
//  Name        string `json:"book_name"`
//  Description string `json:"descript"`
// }
// type DeleteBookResponse struct {
//  BookID string `json:"book_id"`
// }
// var DB *gorm.DB
// var err error
// @Title Create Book
// @Description Creates Books & Returns a Book based on the request
// @Param request body _Book true "Create Book Request"
// @Router /books/create [post]
func CreateNewBookapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	DB.Create(&book)
	json.NewEncoder(w).Encode(book)
}

// @Title Get All Book
// @Description Get All Books based on the request
// @Router /books [get]
func GetAllBooksapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var books []Book
	DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// @Title Get _Book by ID
// @Description Get Books by ID based on the request
// @Param id path int true "Get Books by ID Request"
// @Router /books/{id} [get]
func GetSingleBookapi(w http.ResponseWriter, r *http.Request) { //this will help you to knw single book in an api you can do get request on postman or you can check on booklist
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintln(w, "Please Enter Integer")
	} else {
		var book Book
		DB.First(&book, params["id"])
		json.NewEncoder(w).Encode(book)
	}
}

// @Title Update _Book By ID
// @Description Update Books based on the request
// @Param request body _Book true "Update Book Request"
// @Param id path int true "Update Book Request"
// @Router /books/{id} [patch]
func UpdateExistingBookapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//here we are declaring the err and initializing the err as strconv and under the functions we are passing the params
	_, err := strconv.Atoi(params["id"]) //strconv is an inbuilt package that is used for conversion to and from string representation of basic data types
	if err != nil {                      //strconv is used for string to int type conversion the Atoi is a function with strconv so if you are taking a string and you want to convert in integer then you can use strconv.Atoi
		fmt.Fprintln(w, "Please Enter Integer,please avoid using string") //If the user enters a string then this will be printed that please enter an integer
	} else {
		var book Book
		DB.First(&book, params["id"])
		json.NewDecoder(r.Body).Decode(&book)
		DB.Save(&book)
		json.NewEncoder(w).Encode(book)
	}
}

// @Title Delete Book By ID
// @Description Delete Books based on the request
// @Param id path int true "Delete Book Request"
// @Router /books/{id} [delete]
func DeletePreviousBookapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintln(w, "Please Enter Integer,please avoid using string")
	} else {
		fmt.Println(params)
		var book Book
		DB.Delete(&book, params["id"])
		json.NewEncoder(w).Encode("This Book has been deleted")

	}
}
