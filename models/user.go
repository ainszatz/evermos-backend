package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Phone    string `json:"phone" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"role" gorm:"default:user"` // default user, bisa diubah ke "admin"
	Store    Store  `json:"store" gorm:"foreignKey:UserID"`
}
