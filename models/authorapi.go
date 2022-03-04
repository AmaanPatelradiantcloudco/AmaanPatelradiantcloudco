package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// var validate = validator.New()
/*type Author2 struct {
	gorm.Model
	Name         string `json:"Name"`
	Description  string `json:"description`
	bookswritten string `json:"bookswritten"`
	popularbooks string `json "popularbooks"`
}*/

type Author struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

var validate = validator.New()

func (author *Author) Validate() error {
	err := validate.Struct(author)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}

func CreateAuthorapi(w http.ResponseWriter, r *http.Request) {
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

func GetAuthorsapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var authors []Author
	DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var author Author
	DB.First(&author, params["id"])
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthorapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var author Author
	params := mux.Vars(r)
	DB.First(&author, params["id"])
	json.NewDecoder(r.Body).Decode(&author)
	DB.Save(&author)
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
