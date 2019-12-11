package structs

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Place string
	Total string
}

type Register struct {
	gorm.Model
	First_Name   string
	Last_Name    string
	Username     string
	Password     string
	Email        string
	Phone_Number string
}
