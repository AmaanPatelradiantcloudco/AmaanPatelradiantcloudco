package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

//used in jwt to sign our token
var jwtKey = []byte("secret_key")

var people = map[string]string{
	"user1":       "password1",
	"user2":       "password2",
	"Moinqureshi": "12345",
}

//for passing username and pw from the api

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//to pass the username when the payload is exp
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type User struct {
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

func HashPassword(Password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("unable to genereate hash: %w", err)
	}
	return string(bytes), err
}

func CreateUserapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	var check_user User
	DB.Table("users").Where("username = ?", user.Username).Scan(&check_user)
	fmt.Printf("%#v\n", check_user)
	//func GenerateFromPassword(check_user.password []byte, cost int) ([]byte, error)
	if check_user.Username == "" {
		passd := user.Password
		//	user.Password := models.HashPassword()
		hash, hasherror := HashPassword(passd)
		if hasherror != nil {
			fmt.Println("error hash value not generated", hasherror)
			return
		}

		user.Password = hash
		fmt.Println("Password:", passd)
		fmt.Println("Hash:    ", hash)
		DB.Create(&user)
		fmt.Fprintln(w, "new user created")
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var verify_user User
	DB.Table("users").Where("username = ?", credentials.Username).Scan(&verify_user)
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
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//using claims object and the jwt key, will create a token out of it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//from this token will get a token string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//if no err over here that means we got our token strings over here and we are ready to set those things in our cookies

	//all the details will be stored in our cookie.

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

// Password := "secret"
// hash, _ := models.HashPassword(user.password) // ignore error for the sake of simplicity

// fmt.Println("Password:", Password)
// fmt.Println("Hash:    ", hash)

// match := models.CheckPasswordHash(user.password, hash)
// fmt.Println("Match:   ", match)

// func Password_Check{
// 	var pw_check
// 	func HashPassword(password, string) (string, error) {
// 		bytes, err := bcrypt.GenerateFromPassword([]byte(user.password), 14)
// 		return string(bytes), err
// 	}

// 	func CheckPasswordHash(password, hash string) bool {
// 		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.password))
// 		return err == nil
// 	}

// }

func Homeapi(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		//if no cookie i.e unauthorised access, we want a cookie with a token
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//getting a value of the token and setting the value over here
	tokenStr := cookie.Value

	//parse the claims back with token string and jwtkey
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

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

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
