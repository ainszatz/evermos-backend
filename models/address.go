package models

import "gorm.io/gorm"

type Address struct {
    gorm.Model
    UserID    uint   `json:"user_id"`
    StoreID   *uint  `json:"store_id,omitempty"`
    Address   string `json:"address"`
	Street string `json:"street"`
    City      string `json:"city"`
    Province  string `json:"province"`
    ZipCode   string `json:"zip_code"`
}
