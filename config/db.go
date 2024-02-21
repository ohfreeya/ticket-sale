package config

import (
	"ticketsale/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var models = []interface{}{
	model.User{},
}

func DBInit() error {
	dbConfig := DBConfig()
	var err error

	var dsn = dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ")/" + dbConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	// fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return nil
}

func DBMigrate() {
	DB.AutoMigrate(models...)
}
