package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Title    string `json:"title" binding:"required,max=50"`
	Abstract string `json:"abstract" binding:"required,max=255"`
	FileName string `json:"file_name" binding:"required,max=255"`
	UserId   int    `json:"user_id"`
	User     User
	Comment  []Comment
}

func (p Paper) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        uint
		Title     string `json:"title"`
		Abstract  string `json:"abstract" binding:"required,max=255"`
		FileName  string `json:"file_name" binding:"required,max=255"`
		UserId    int    `json:"user_id"`
		User      User
		CreatedAt string
	}{
		Title:     p.Title,
		Abstract:  p.Abstract,
		FileName:  p.FileName,
		UserId:    p.UserId,
		User:      p.User,
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
