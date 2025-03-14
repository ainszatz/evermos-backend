package models

import "gorm.io/gorm"

type ProductLog struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Change    int     `json:"change"` // Bisa negatif (penjualan) atau positif (restock)
	Note      string  `json:"note"`
}
