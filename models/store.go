package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Name   string `json:"name" gorm:"not null"`
}
