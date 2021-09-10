package model

import (
    "github.com/jinzhu/gorm"
)

type Comment struct {
    gorm.Model
    Name string
	Body string
	PaperId int
}