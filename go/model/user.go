package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" binding:"required,max=50"`
	Email     string `json:"email" binding:"required,email,max=255"`
	Password  string `json:"password" binding:"required,min=8,max=16"`
	AdminRole bool   `json:"admin_role"`
	Paper     []Paper
}
