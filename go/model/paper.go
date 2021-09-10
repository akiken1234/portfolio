package model

import (
    "github.com/jinzhu/gorm"
)

type Paper struct {
    gorm.Model
    Title string
	Abstract string  
	FileName string `gorm:"unique"`
	UserId int
}