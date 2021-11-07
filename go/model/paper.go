package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Title    string `json:"title" validate:"required,max=50" gorm:"not null;size:50"`
	Abstract string `json:"abstract" validate:"required,max=200" gorm:"not null;size:200"`
	FileName string `json:"file_name" validate:"required" gorm:"not null;unique"`
	UserId   int    `json:"user_id" validate:"required" gorm:"not null"`
	User     User
}

func (p Paper) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        uint
		Title     string `json:"title" validate:"required,max=50" gorm:"not null;size:50"`
		Abstract  string `json:"abstract" validate:"required,max=200" gorm:"not null;size:200"`
		FileName  string `json:"file_name" validate:"required" gorm:"not null"`
		UserId    int    `json:"user_id" validate:"required" gorm:"not null"`
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
