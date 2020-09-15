package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

	db, err := gorm.Open("postgres", "user=sairambagam password=gorm dbname=test sslmode=disable")

	if err != nil {
		panic(err)
	}

	fmt.Println("db is ", db.Value)
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	// create a table if it doesnt exist already
	db.CreateTable(&Owner{})

	owner := Owner{
		FirstName: "Sai",
		LastName:  "Bagam",
	}
	db.Create(&owner)
}

type Owner struct {
	gorm.Model
	FirstName string
	LastName  string
}
