package controller

import (
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Comment struct{}

func NewComment() *Comment {
	return &Comment{}
}

func (t *Comment) List(c *gin.Context) {
	db := db.DB()

	var comments []model.Comment
	err := db.Find(&comments).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, comments)
	}
}

func (t *Comment) Get(c *gin.Context) {
	db := db.DB()

	var comment model.Comment
	err := db.First(&comment, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, comment)
	}
}

func (t *Comment) Create(c *gin.Context) {
	db := db.DB()

	var comment model.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = db.Create(&comment).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

func (t *Comment) Delete(c *gin.Context) {
	db := db.DB()

	var comment model.Comment
	err := db.Delete(&comment, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
