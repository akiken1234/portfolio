package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `validate:"required,max=50"`
	Email     string `validate:"email,required,unique,max=255"`
	Password  string `validate:"required,min=8,max=16"`
	AdminRole bool
	Paper     []Paper
}
