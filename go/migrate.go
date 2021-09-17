package main

import (
	"local.packages/db"
	"local.packages/model"
)

func main() {
	db.Connection()
	defer db.DB().Close()

	db.DB().AutoMigrate(&model.User{})
	db.DB().AutoMigrate(&model.Paper{})
	db.DB().AutoMigrate(&model.Message{})
	db.DB().AutoMigrate(&model.Comment{})
}
