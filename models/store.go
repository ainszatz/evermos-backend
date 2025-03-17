package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Name   string `json:"name" gorm:"not null"`
	LogoURL string `json:"logo_url"`
}
