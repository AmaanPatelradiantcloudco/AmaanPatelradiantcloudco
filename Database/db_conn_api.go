package Database

import (
	"fmt"            //Format packager which is used for printf statements
	"restapi/models" //here we are importing models which have different author,books and users.go file

	"gorm.io/driver/mysql" //supports to write a driver here we are taking mysql
	"gorm.io/gorm"         //used for crud operations and implementation of initialize router
)

var DB *gorm.DB //we are declaring a new variable Db it is taking the gorm package which is for intialmigration and DB struct
var err error   //here we are declaring the err variable and error interface

const DNS = "root:@tcp(127.0.0.1:3306)/phpmyadmin?parseTime=true" //this is domain name system here we are simply giving localhost
//3306 which is simply there on the xampp and 127.0.0.1 refers to localhost and the database is phpmyadmin

func InitialMigration() { //here we are creating function initialmigration which is taken from gorm
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{}) //DB is variable and after there's error variable,then we have gorm.gorm.Open creates a new connection pool every time,we have config which is nothing but config struct
	if err != nil {                                      //if err not equal to nil
		fmt.Println(err.Error())          //this is called from builtin.go
		panic("Cannot Connect to the DB") //another print statement
		println("Please check the connection")

	}
	DB.AutoMigrate(&models.Authorapi{}, &models.Book{}, &models.User{}) //here we are taking the automigrate which is there inbuilt in
	//migrator.go-this is from gorm package
}

// func createRandomUser(t *testing.T) User {
// 	hashedPassword, err :=HashPassword(RandomString(6))
// 	require.NoError(t, err)
// 	arg :=  User{
// 		HashedPassword: hashedPassword,
// 	}
// }

func GetDB() *gorm.DB {
	return DB
}
