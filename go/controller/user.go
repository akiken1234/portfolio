package controller

import (
	"local.packages/db"
	"local.packages/model"

	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (t *User) List(c *gin.Context) {
	db := db.DB()

	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, users)
	}
}

func (t *User) Get(c *gin.Context) {
	db := db.DB()

	var user model.User
	err := db.First(&user, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, user)
	}
}

func (t *User) Create(c *gin.Context) {
	db := db.DB()

	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = db.Create(&user).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

func (t *User) Update(c *gin.Context) {
	db := db.DB()

	var user model.User
	err := db.First(&user, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "user not found")
		return
	}
	err = c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = db.Save(&user).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

func (t *User) Delete(c *gin.Context) {
	db := db.DB()

	var user model.User
	err := db.Delete(&user, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}
