package controller

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-gonic/gin"
)

type Paper struct{}

func NewPaper() *Paper {
	return &Paper{}
}

func (t *Paper) List(c *gin.Context) {
	db := db.DB()
	log.Println(c.Param("list"))

	var papers []model.Paper
	err := db.Preload("User").Find(&papers).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, papers)
	}
}

func (t *Paper) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("./files/" + c.Request.FormValue("file_name"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
	}
	content_type := http.DetectContentType(data[:512])

	c.Data(http.StatusOK, content_type, data)
}

func (t *Paper) Create(c *gin.Context) {
	db := db.DB()

	title := c.Request.FormValue("title")
	abstract := c.Request.FormValue("abstract")
	file_name := c.Request.FormValue("file_name")
	paper := model.Paper{Title: title, Abstract: abstract, FileName: file_name}
	err := db.Create(&paper).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
	// pdfファイル保存
	pdf, header, _ := c.Request.FormFile("file")
	saveFile, _ := os.Create("./files/" + header.Filename)
	defer saveFile.Close()
	io.Copy(saveFile, pdf)
}

func (t *Paper) Delete(c *gin.Context) {
	db := db.DB()

	var paper model.Paper
	err := db.Delete(&paper, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
