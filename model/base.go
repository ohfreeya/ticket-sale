package model

import "ticketsale/config"

func Migrate() {
	var models = []interface{}{
		&User{},
	}
	config.DB.AutoMigrate(models...)
}
