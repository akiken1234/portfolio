package model

import (
	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Title    string `json:"title" binding:"required,max=50"`
	Abstract string `json:"abstract" binding:"required,max=255"`
	FileName string `json:"file_name" binding:"required,max=255"`
	UserId   int    `json:"user_id"`
	Comment  []Comment
}
