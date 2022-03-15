package models

import (
	"encoding/json" //package encoding/json
	"fmt"           //package fmt this is format package which is used for print statements
	"net/http"      //this package is for http request specifically GET request
	"strconv"       //strconv which is equivalent

	"github.com/go-playground/validator/v10" //used for implementing value validations for struct
	"github.com/gorilla/mux"                 //this is used for implementing request router and for matching incoming request to their respective handler
	"gorm.io/gorm"                           //gorm provides crud operation and can also be used for initial migration and  creation of database schema
)

var DB *gorm.DB //we are declaring two variable of db over here so that we can access them here
var err error   //same goes for this one too

type Authorapi struct { //we are declaring the struct which is similar to classes in languages
	gorm.Model
	Name        string `json:"name"`        //for input we are taking name we are declaring in the json format
	Description string `json:"description"` ///for input we are taking description we are declaring in the json format
}

var validate = validator.New() //here  we are declaring the var validate and creating an instance validator.New()

func (author *Authorapi) Validate() error { //we are declaring the func author which has author variable and then Author struct and then calling Validate()
	err := validate.Struct(author) //here we are declaring and initializing err which is var validate.Struct and author variable as argument
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok { //if condition if there is invalidation error is there returns err
			return err
		}
		return err
	}
	return nil
}

func CreateNewAuthorapi(w http.ResponseWriter, r *http.Request) { //simple reesponse writer
	w.Header().Set("Content-Type", "application/json") //we are setting the header which have content-type,application/json
	var authorapi Authorapi                            //declaring variabble author over here and then the struct Author
	json.NewDecoder(r.Body).Decode(&authorapi)         //we are giving json decoder, r is for http request then we are decoding the variable author
	var check_author Authorapi                         //new variable check_author
	DB.Table("authors").Where("name = ?", authorapi.Name).Scan(&check_author)
	//DB.Table("users").Where("username=?, password=?",user.username, user.password).
	fmt.Printf("%#v\n", check_author)

	if check_author.Name == "" {
		DB.Create(&authorapi)
		fmt.Fprintln(w, "new Author created")
		json.NewEncoder(w).Encode(authorapi)
	} else {
		fmt.Fprintln(w, "Author already exist")
	}

}

func GetAllAuthorsapi(w http.ResponseWriter, r *http.Request) { //simple response writer
	w.Header().Set("Content-Type", "application/json") //setting the headers
	var authorsall []Authorapi                         //var authors for getting all the authors
	DB.Find(&authorsall)                               //here we are using the variable DB which is there in database file and we are finding the author in database using find which is again inbuilt and there in finisher_api.go
	json.NewEncoder(w).Encode(authorsall)              //encoding the variable author
}

func GetSingleAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, err2 := strconv.Atoi(params["id"]) //declaring and initializing the err2 and we are using strconv this function with Atoi is used for string to integer conversions
	//here we are checking into condition so if the error is not nil it will printed as following
	if err2 != nil {
		fmt.Fprintln(w, "Please Enter Integer,please avoid using string", err2)
	} else { //we have declared var params,and we have initialize mux router which can be
		var authorsingle Authorapi              //done by using go get (package name)
		DB.First(&authorsingle, params["id"])   //we are using the variable author over here and then for the params since we are
		json.NewEncoder(w).Encode(authorsingle) //declaring the trying to get single author using id
	}
}

func UpdateAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var author Authorapi
	params := mux.Vars(r)
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintln(w, "Please Enter Integer")
	} else {
		DB.First(&author, params["id"])
		json.NewDecoder(r.Body).Decode(&author)
		DB.Save(&author)
		json.NewEncoder(w).Encode(author)
	}

}

func DeleteAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var author Authorapi
	DB.Delete(&author, params["id"])
	fmt.Fprintf(w, "Successfully deleted ")
}
