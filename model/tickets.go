package model

import (
	"time"

	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model
	OwnerID   int
	Count     int
	Introduce string `gorm:"type:text"`
	ExpiresAt time.Time
	User      User `gorm:"foreignKey:OwnerID"`
}

type TicketsType struct {
	gorm.Model
	Name      string
	Introduce string `gorm:"type:text"`
}

type TicketsSalePrice struct {
	gorm.Model
	TicketsID     int
	TicketsTypeID int
	Price         float64
	ExpiresAt     time.Time
	Tickets       Tickets `gorm:"foreignKey:TicketsID"`
	TicketsType   TicketsType `gorm:"foreignKey:TicketsTypeID"`
}
