package model

import (
	"ticketsale/config"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID   uint    `gorm:"comment: 使用者ID"`
	TicketID uint    `gorm:"comment: 票券ID"`
	Status   int8    `gorm:"comment: 訂單狀態"`
	Count    int     `gorm:"comment: 購買數量"`
	Price    float64 `gorm:"comment: 購買價格"`
	Total    float64 `gorm:"comment: 總價"`
	Coupon   string  `gorm:"comment: 折價券"`
	User     User    `gorm:"foreignKey:UserID"`
	Tickets  Tickets `gorm:"foreignKey:TicketID"`
}

// order's method
func (o *Orders) Find(params map[string]interface{}) (ret Orders, err error) {
	result := config.DB.Where(params).First(&ret)
	if result.RowsAffected == 0 {
		err = result.Error
	}
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (o *Orders) Create() error {
	err := config.DB.Create(o)
	return err.Error
}

func (o *Orders) Update() error {
	err := config.DB.Save(o)
	return err.Error
}

func (o *Orders) Delete() error {
	err := config.DB.Delete(&o)
	return err.Error
}
