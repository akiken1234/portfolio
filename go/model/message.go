package model

import (
    "github.com/jinzhu/gorm"
)

type Message struct {
    gorm.Model
	Name string
    Title string
	Body string
}