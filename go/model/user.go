package model

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Name string
	Email string `gorm:"unique"`
	Password string
	Paper []Paper
	AdminRole bool
}