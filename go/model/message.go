package model

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Name       string `json:"name" binding:"required,max=50"`
	Title      string `json:"title" binding:"required,max=50"`
	Body       string `json:"body" binding:"required,max=255"`
	FromUserId int    `json:"from_user_id" binding:"required,max=50"`
	ToUserId   int    `json:"to_user_id" binding:"required,max=50"`
}
