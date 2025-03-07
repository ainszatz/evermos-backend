package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	StoreID    uint    `json:"store_id"`
	CategoryID uint    `json:"category_id"`
	Name       string  `json:"name" gorm:"not null"`
	Price      float64 `json:"price" gorm:"not null"`
	Stock      int     `json:"stock" gorm:"not null"`
}
