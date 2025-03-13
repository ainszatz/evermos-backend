package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `json:"name" gorm:"not null"`
	Price   float64 `json:"price" gorm:"not null"`
	Stock   int     `json:"stock" gorm:"not null"`
	StoreID uint    `json:"store_id" gorm:"not null;index"`
	Store   Store   `json:"store" gorm:"foreignKey:StoreID"`
}
