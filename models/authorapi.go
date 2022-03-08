package models

import (
	"encoding/json" //package encoding/json
	"fmt"           //package fmt this is format package which is used for print statements
	"net/http"      //this package is for http request specifically GET request

	"github.com/go-playground/validator/v10" //used for implementing value validations for struct
	"github.com/gorilla/mux"                 //this is used for implementing request router and for matching incoming request to their respective handler
	"gorm.io/gorm"                           //gorm provides crud operation and can also be used for initial migration and  creation of database schema
)

var DB *gorm.DB //we are declaring two variable of db over here so that we can access them here
var err error   //same goes for this one too

// var validate = validator.New()
/*type Author2 struct {
	gorm.Model
	Name         string `json:"Name"`
	Description  string `json:"description`
	bookswritten string `json:"bookswritten"`
	popularbooks string `json "popularbooks"`
}*/

type Author struct { //we are declaring the struct which is similar to classes in languages
	gorm.Model
	Name        string `json:"name"`        //for input we are taking name we are declaring in the json format
	Description string `json:"description"` ///for input we are taking description we are declaring in the json format
}

var validate = validator.New() //here  we are declaring the var validate and creating an instance validator.New()

func (author *Author) Validate() error { //we are declaring the func author which has author variable and then Author struct and then calling Validate()
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
	var author Author                                  //declaring variabble author over here and then the struct Author
	json.NewDecoder(r.Body).Decode(&author)            //we are giving json decoder, r is for http request then we are decoding the variable author
	var check_author Author                            //new variable check_author
	DB.Table("authors").Where("name = ?", author.Name).Scan(&check_author)
	//DB.Table("users").Where("username=?, password=?",user.username, user.password).
	fmt.Printf("%#v\n", check_author)

	if check_author.Name == "" {
		DB.Create(&author)
		fmt.Fprintln(w, "new user created")
		json.NewEncoder(w).Encode(author)
	} else {
		fmt.Fprintln(w, "user already exist")
	}

	// type CreateAuthor struct {
	//  Name        string `json:"book_name"`
	//  Description string `json:"descript"`
	// }
	// type GetAuthorsRequest struct {
	//  BookID string `json:"book_id"`
	// }
	// type AuthorsResponse struct {
	//  BookID string `json:"book_id"`
	// }
	// type GetAuthorRequest struct {
	//  Name        string `json:"book_name"`
	//  Description string `json:"descript"`
	// }
	// type GetAuthorResponse struct {
	//  BookID string `json:"book_id"`
	// }
	// type UpdateAuthorRequest struct {
	//  Name        string `json:"book_name"`
	//  Description string `json:"descript"`
	// }
	// type UpdateAuthorResponse struct {
	//  BookID string `json:"book_id"`
	// }
	// type DeleteAuthorRequest struct {
	//  Name        string `json:"book_name"`
	//  Description string `json:"descript"`
	// }
	// type DeleteAuthorResponse struct {
	//  BookID string `json:"book_id"`
	// }
	// var DB *gorm.DB
	// var err error
	// @Title Create Book
	// @Description Creates Books & Returns a Book based on the request
	// @Param request body _Book true "Create Book Request"
	// @Router /books/create [post]

	//	json.NewEncoder(w).Encode(author)

	// if err != nil {
	// 		panic(err)
	// 	 }

	// fmt.Printf("%#v\n", DB.Where(author.Name).Take(&Author{}))
	// if err != nil{
	// 	if DB.Where(author.Name).Take(Author{}).Error != nil {
	// 	fmt.Fprintln(w, "data already there")
	// 	}
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", DB.Where(author.Name).Take(&Author{}).Error)
	// fmt.Printf("%#v\n", DB.Where(author.Name).Take(&Author{}))
	// if DB.Where(author.Name).Take(&Author{}).Error != nil {
	// // }
	//	DB.Create(&author)

}

func GetAllAuthorsapi(w http.ResponseWriter, r *http.Request) { //simple response writer
	w.Header().Set("Content-Type", "applciation/json") //setting the headers
	var authors []Author                               //var authors for getting all the authors
	DB.Find(&authors)                                  //here we are using the variable DB which is there in database file and we are finding the author in database using find which is again inbuilt and there in finisher_api.go
	json.NewEncoder(w).Encode(authors)                 //encoding the variable author
}

func GetAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)             //we have declared var params,and we have initialize mux router which can be
	var author Author                 //done by using go get (package name)
	DB.First(&author, params["id"])   //we are using the variable author over here and then for the params since we are
	json.NewEncoder(w).Encode(author) //declaring the trying to get single author using id
}

func UpdateAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var author Author
	params := mux.Vars(r)                   //declarings the params variable since we need to pass the param which is id over here
	DB.First(&author, params["id"])         //we are using db.first so it goes to author variable and the id of the author
	json.NewDecoder(r.Body).Decode(&author) //decoding the author
	DB.Save(&author)                        //saving the new author that is updated and posted using post method
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var author Author
	DB.Delete(&author, params["id"])
	json.NewEncoder(w).Encode("The Author Has been successfully deleted")
}

// const DNS = "root:@tcp(127.0.0.1:3306)/phpmyadmin?parseTime=true"

// func init() {
// 	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot Connect TO the Database")
// 	}
// 	models.InitialMigration()
// 	DB = models.GetDB()
// 	DB.AutoMigrate(&Author{})
// }

// func (author *Author) validate() error {
// 	err := validate.Struct(author)
// 	if err != nil {
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			return err
// 		}
// 		return err
// 	}
// 	return nil
// }
func CreateAuthor2api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author Author
	json.NewDecoder(r.Body).Decode(&author)
	var check_author Author
	DB.Table("authors").Where("name = ?", author.Name).Scan(&check_author)
	//DB.Table("users").Where("username=?, password=?",user.username, user.password).
	fmt.Printf("%#v\n", check_author)

	if check_author.Name == "" {
		DB.Create(&author)
		fmt.Fprintln(w, "new user created")
		json.NewEncoder(w).Encode(author)
	} else {
		fmt.Fprintln(w, "user already exist")
	}

	//	json.NewEncoder(w).Encode(author)

	// if err != nil {
	// 		panic(err)
	// 	 }

	// fmt.Printf("%#v\n", DB.Where(author.Name).Take(&Author{}))
	// if err != nil{
	// 	if DB.Where(author.Name).Take(Author{}).Error != nil {
	// 	fmt.Fprintln(w, "data already there")
	// 	}
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", DB.Where(author.Name).Take(&Author{}).Error)
	// fmt.Printf("%#v\n", DB.Where(author.Name).Take(&Author{}))
	// if DB.Where(author.Name).Take(&Author{}).Error != nil {
	// // }
	//	DB.Create(&author)

}

/*
func GetAuthors2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var authors []Author
	DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func GetAuthor2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var author Author
	DB.First(&author, params["id"])
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var author Author
	params := mux.Vars(r)
	DB.First(&author, params["id"])
	json.NewDecoder(r.Body).Decode(&author)
	DB.Save(&author)
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var author Author
	DB.Delete(&author, params["id"])
	json.NewEncoder(w).Encode("The Author Has been successfully deleted")
}
*/
