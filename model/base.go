package model

import "ticketsale/config"

func Migrate() {
	var models = []interface{}{
		&User{},
		&Tickets{},
		&TicketsType{},
		&TicketsSalePrice{},
	}
	config.DB.AutoMigrate(models...)
}
