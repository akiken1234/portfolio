package controller

import (
	"fmt"
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

func (t *Paper) Download(c *gin.Context) {
	log.Println(c.Request.FormValue("file_name"))
	data, err := ioutil.ReadFile("./files/" + c.Request.FormValue("file_name"))
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file_1": data,
		"file_2": string(data),
	})
}

// dst := fmt.Sprintf("%s/%s", "/go/src/github.com/akimotokensaku/portfolio/go/files", c.Request.FormValue("file_name"))
	// log.Println(dst)
	// b, err := ioutil.ReadFile(dst)
	// if err != nil {
	// 	log.Println(222)
	// 	c.String(http.StatusInternalServerError, "Server Error")
	// }
	// m := http.DetectContentType(b[:512])
	// c.JSON(http.StatusOK, gin.H{
	// 	m: b,
	// })

// func (t *Paper) Download(c *gin.Context) {
// 	log.Println(333)
// 	dst := fmt.Sprintf("%s/%s", "./files/", c.Request.FormValue("file_name"))
// 	pdf, err := ioutil.ReadFile(dst)
// 	if err != nil {
// 		log.Println(222)
// 		c.String(http.StatusInternalServerError, "Server Error")
// 	}
// 	log.Println(444)

// 	c.JSON(200, gin.H{"file": pdf})

// 	// m := http.DetectContentType(b[:512])
// 	// return m, b, nil
// }

func (t *Paper) Delete(c *gin.Context) {
	db := db.DB()

	var paper model.Paper
	err := db.Delete(&paper, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
