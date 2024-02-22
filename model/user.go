package model

import (
	"ticketsale/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Account  string
	Password string
	Email    string `gorm:"unique"`
	Type     int8   `gorm:"default:0"`
	Phone    string
	Sex      int8 `gorm:"default:0"`
}

func (user *User) Find(params map[string]interface{}) (ret User, err error){
	result := config.DB.Where(params).First(&ret)
	if result.RowsAffected == 0 {
		err = result.Error
	}
	if result.Error != nil {
		err = result.Error
	}
	return 
}

func (user *User) Create() error {
	err := config.DB.Create(user)
	return err.Error
}

func (user *User) CheckEmailExist() bool {
	var u User
	config.DB.Where("email = ?", user.Email).First(&u)
	return u.ID != 0
}
