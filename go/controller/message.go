package controller

import (
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct{}

func NewMessage() *Message {
	return &Message{}
}

func (t *Message) List(c *gin.Context) {
	db := db.DB()

	var messages []model.Message
	err := db.Find(&messages).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, messages)
	}
}

func (t *Message) Get(c *gin.Context) {
	db := db.DB()

	var message model.Message
	err := db.First(&message, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, message)
	}
}

func (t *Message) Create(c *gin.Context) {
	db := db.DB()

	var message model.Message
	err := c.BindJSON(&message)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = db.Create(&message).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

func (t *Message) Delete(c *gin.Context) {
	db := db.DB()

	var message model.Message
	err := db.Delete(&message, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
