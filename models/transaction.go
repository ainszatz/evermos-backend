package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	User       User    `json:"user" gorm:"foreignKey:UserID"`
	ProductID  uint    `json:"product_id"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
