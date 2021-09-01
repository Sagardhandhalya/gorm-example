package main

import (
	"fmt"

	"gorm.io/gorm"
)

func init() {
	ConnectToDb()
}

func ShowOutput(tx *gorm.DB) {
	if tx.Error != nil {
		fmt.Print(tx.Error.Error())
	} else {
		fmt.Printf("Success - Opration Success %d row affected ___!! \n\n", tx.RowsAffected)
	}
}

func main() {
	// // creating table in Db
	Db.AutoMigrate(&User{}, &CreditCard{}, &Car{}, &Language{})

	// creating new user
	user := User{
		Name:    "samir",
		Age:     99,
		IsAdmin: false,
		CreditCard: CreditCard{
			CardNo: "23232242454",
		},
		Languages: []Language{{Name: "C++"}},
		Cars:      []Car{{Name: "xyz"}, {Name: "Hundai"}},
	}
	Db.Model(User{}).Create(&user)

	/*
	   LET"S QUERY RELATIONS
	*/

	// 1. fetch first user
	// usr := User{}
	// ShowOutput(Db.First(&usr))

	// 2. fetch all car of the user

}
