package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name    string   `json:"name" gorm:"not null"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}
