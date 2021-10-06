package controller

import (
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Paper struct{}

func NewPaper() *Paper {
	return &Paper{}
}

func (t *Paper) List(c *gin.Context) {
	db := db.DB()

	var papers []model.Paper
	err := db.Preload("User").Find(&papers).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, papers)
	}
}

func (t *Paper) Get(c *gin.Context) {
	db := db.DB()

	var paper model.Paper
	err := db.First(&paper, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, paper)
	}
}

func (t *Paper) Create(c *gin.Context) {
	db := db.DB()

	var paper model.Paper
	err := c.BindJSON(&paper)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = db.Create(&paper).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

func (t *Paper) Delete(c *gin.Context) {
	db := db.DB()

	var paper model.Paper
	err := db.Delete(&paper, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
