package models

import (
	"encoding/json" //for encoding and decoding package
	"fmt"           //for print statements
	"net/http"      //for getting GET request
	"time"          //provides fucnctionality of time

	"github.com/gorilla/mux"     //for matching request
	"golang.org/x/crypto/bcrypt" //package for operations of encrypting etc.

	"github.com/dgrijalva/jwt-go" //package for jwt-token,standarclaims
	"gorm.io/gorm"                //for crud operations and initialmigration
)

//used in jwt to sign our token
var jwtKey = []byte("secret_key")

/*var people = map[string]string{
	"user1":      "password1",
	"user2":      "password2",
	"amaanpatel": "23460728",
	"amaanpatel7868":"6846t8425"  //this username and password that is given is to check after posting are they there or not without db connection
}*/

//for passing username and pw from the api

type Credentials struct {
	Email    string `json:"email"` //taking the credentials in json format
	Password string `json:"password"`
}

//to pass the username when the payload is exp
type Claims struct {
	Email              string `json:"email"`
	jwt.StandardClaims        //standardclaims from jwt library
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Usernew struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}
func GetUserseparate(w http.ResponseWriter, r *http.Request) { //func is created and we have response writer
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)           //declaring and initializing the params as we did for other funcs
	var user User                   //declaring variable user and then struct User
	DB.First(&user, params["id"])   //here for db we are passing the var name and id
	json.NewEncoder(w).Encode(user) //we are here encoding the variable user
}
func HashPassword(Password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost) //generatefrompasswd is from bcrypt.go,variable passwd
	if err != nil {
		return "", fmt.Errorf("unable to genereate hash: %w", err)
	}
	return string(bytes), err
}

func CreateNewUserapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	var check_user User
	DB.Table("users").Where("username = ?", user.Username).Scan(&check_user)
	fmt.Printf("%#v\n", check_user)
	//func GenerateFromPassword(check_user.password []byte, cost int) ([]byte, error)
	if check_user.Username == "" {
		passd := user.Password //declaring passd variable,user variable and then calling password
		//	user.Password := models.HashPassword()
		hash, hasherror := HashPassword(passd)
		if hasherror != nil {
			fmt.Println("error hash value not generated", hasherror)
			return
		}

		user.Password = hash
		fmt.Println("Password:", passd) //print for password
		fmt.Println("Hash:    ", hash)  //print for hash
		DB.Create(&user)
		fmt.Fprintln(w, "new user created") //print statement
		json.NewEncoder(w).Encode(user)
	} else {
		fmt.Fprintln(w, "user already exist")
	}

}

func CheckPasswordHash(Password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	return err == nil
}

func Loginapi(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil { //if error is not equal to nil
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//expectedPassword,ok :=users[credentials.Username]
	//	if !ok || expectedPassword != credentials.Password{

	//}//
	var verify_user User
	DB.Table("users").Where("username = ?", credentials.Email).Scan(&verify_user)
	fmt.Printf("%#v\n", verify_user)
	//if data not available in the map or the password did not match then throw error
	ispwd := CheckPasswordHash(credentials.Password, verify_user.Password)
	fmt.Println("Match:   ", err)

	if !ispwd {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if !ok || expectedPassword != credentials.Password {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	expirationTime := time.Now().Add(time.Minute * 5) //we are setting the expiration time to 5 mins
	claims := &Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //calling the expiartion time,unix is taken from inbuilt time.go
		},
	}

	//using claims object and the jwt key, will create a token out of it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //we are using hs256 method there are lot of available
	//from this token will get a token string and we are passing the jwtkey which we have given above
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //condition for internal server error
		return
	}
	//if no err over here that means we got our token strings over here and we are ready to set those things in our cookies

	//all the details will be stored in our cookie.,the data will be stored in cookie

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",        //name will be token for cookie
			Value:   tokenString,    //value is tokenstring that we will get from token
			Expires: expirationTime, //taken from the expiration time that we declared above
		})

}

func HomePageapi(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token") //cookie variable ,err variable,we need to get the token cookie from here
	if err != nil {
		//if no cookie i.e unauthorised access, we want a cookie with a token
		if err == http.ErrNoCookie { //if there is no cookie then we will do unauthorized access
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//getting a value of the token and setting the value over here
	tokenStr := cookie.Value //we are getting the tokenstring from the cookie

	//parse the claims back with token string and jwtkey
	claims := &Claims{} //reference of the claims

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, //parsing the tokenstring,claims and function which will return jwtkey and nil
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil { //condition for jwt
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid { //if tkn condition
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Email))) //we are giving welcome when it is authenicated

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}
