package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Name    string `validate:"required,unique,max=50"`
	Body    string `validate:"required,unique,max=255"`
	PaperId int
}
