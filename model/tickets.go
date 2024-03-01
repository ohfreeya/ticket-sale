package model

import (
	"ticketsale/config"
	"time"

	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model
	OwnerID   int
	Name      string `gorm:"unique"`
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

// ticket's method
func (t *Tickets) Find(params map[string]interface{}) (ret Tickets, err error) {
	result := config.DB.Where(params).First(&ret)
	if result.RowsAffected == 0 {
		err = result.Error
	}
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (t *Tickets) Create() error {
	err := config.DB.Create(t)
	return err.Error
}

func (t *Tickets) Update() error {
	err := config.DB.Save(t)
	return err.Error
}

func (t *Tickets) Delete() error {
	err := config.DB.Delete(t)
	return err.Error
}

// tickets type's method
func (t *TicketsType) Create() error {
	err := config.DB.Create(t)
	return err.Error
}

func (t *TicketsType) Update() error {
	err := config.DB.Save(t)
	return err.Error
}

func (t *TicketsType) Delete() error {
	err := config.DB.Delete(t)
	return err.Error
}

// tickets sale price's method
func (t *TicketsSalePrice) Create() error {
	err := config.DB.Create(t)
	return err.Error
}

func (t *TicketsSalePrice) Update() error {
	err := config.DB.Save(t)
	return err.Error
}

func (t *TicketsSalePrice) Delete() error {
	err := config.DB.Delete(t)
	return err.Error
}

