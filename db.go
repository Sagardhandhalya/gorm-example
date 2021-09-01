package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectToDb() {

	if Db == nil {
		Db, err = gorm.Open(mysql.Open("root:9081606040@tcp(127.0.0.1:3306)/gorm_test?parseTime=true"), &gorm.Config{})

		if err != nil {
			log.Fatal("not able to connect to database ")
			panic(err.Error())
		} else {

			fmt.Println("conneted to databse -->")
		}
	}
}
