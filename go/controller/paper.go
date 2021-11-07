package controller

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-gonic/gin"
)

type Paper struct{}

func NewPaper() *Paper {
	return &Paper{}
}

// 論文全件取得
func (t *Paper) List(c *gin.Context) {
	db := db.DB()

	var papers []model.Paper
	if err := db.Preload("User").Find(&papers).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, papers)
	}
}

// ログインユーザーに紐づく論文取得
func (t *Paper) Get(c *gin.Context) {
	db := db.DB()

	var papers []model.Paper
	login_id := c.Request.FormValue("id")
	if err := db.Preload("User").Find(&papers, `user_id = ?`, login_id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, papers)
	}
}

// 論文作成
func (t *Paper) Create(c *gin.Context) {
	db := db.DB()

	title := c.Request.FormValue("title")
	abstract := c.Request.FormValue("abstract")
	file_name := c.Request.FormValue("file_name")
	user_id, err := strconv.Atoi(c.Request.FormValue("user_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}

	paper := model.Paper{Title: title, Abstract: abstract, FileName: file_name, UserId: user_id}

	if err := db.Create(&paper).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
	// pdfファイル保存
	pdf, header, _ := c.Request.FormFile("file")
	saveFile, _ := os.Create("./storage/" + header.Filename)
	defer saveFile.Close()
	io.Copy(saveFile, pdf)
}

// 論文ダウンロード
func (t *Paper) Download(c *gin.Context) {
	data, err := ioutil.ReadFile("./storage/" + c.Request.FormValue("file_name"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
	}
	content_type := http.DetectContentType(data[:512])

	c.Data(http.StatusOK, content_type, data)
}

// 論文削除
func (t *Paper) Delete(c *gin.Context) {
	db := db.DB()

	var paper model.Paper
	if err := db.Where(`file_name = ?`, c.Query("file_name")).Delete(&paper).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
