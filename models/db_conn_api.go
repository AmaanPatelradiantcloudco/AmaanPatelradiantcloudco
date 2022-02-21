package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:@tcp(127.0.0.1:3306)/phpmyadmin?parseTime=true"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot Connect to the DB")
	}
	DB.AutoMigrate(&Author{}, &Book{}, &User{})
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
