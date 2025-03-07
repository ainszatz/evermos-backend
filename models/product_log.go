package models

import "gorm.io/gorm"

type ProductLog struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
}
