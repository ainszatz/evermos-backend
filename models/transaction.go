package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID     uint          `json:"user_id"`
	TotalPrice float64       `json:"total_price"`
	Status     string        `json:"status" gorm:"default:'pending'"`
	Logs       []ProductLog `json:"logs" gorm:"foreignKey:TransactionID"`
}
