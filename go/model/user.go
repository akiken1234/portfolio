package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" validate:"required,max=20" gorm:"not null;size:20"`
	Email     string `json:"email" validate:"required,email" gorm:"not null;unique"`
	Password  string `json:"password" validate:"required,min=8,max=32" gorm:"not null"`
	AdminRole bool   `json:"admin_role" gorm:"default:0"`
	Paper     []Paper
}
