package controller

import (
	"log"
	"net/http"

	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	log.Println(333)
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	hash_password, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	user := model.User{Name: name, Email: email, Password: hash_password}

	err := db.Create(&user).Error
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
