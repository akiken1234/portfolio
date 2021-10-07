package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Name    string `json:"name" binding:"required,max=50"`
	Body    string `json:"body" binding:"required,max=255"`
	PaperId int    `json:"paper_id"`
	Paper Paper
}
