package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	PostalCode  string `json:"postal_code"`
	Detail      string `json:"detail"`
}
