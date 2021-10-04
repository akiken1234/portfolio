package main

import (
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
)

func main() {
	db.Connection()
	defer db.DB().Close()

	db.DB().AutoMigrate(&model.User{})
	db.DB().AutoMigrate(&model.Paper{})
	db.DB().AutoMigrate(&model.Message{})
	db.DB().AutoMigrate(&model.Comment{})
}
