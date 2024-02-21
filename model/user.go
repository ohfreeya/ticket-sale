package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string 
	Account string
	Password string
	Email string `gorm:"unique"`
	Type int8 `gorm:"default:0"`
	Phone string
	Sex int8 `gorm:"default:0"`
}