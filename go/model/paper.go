package model

import (
	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Title    string `validate:"required,unique,max=50"`
	Abstract string `validate:"required,unique,max=255"`
	FileName string `validate:"required,unique,max=255"`
	UserId   int
	Comment  []Comment
}
