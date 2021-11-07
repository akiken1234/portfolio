package controller

import (
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

// ユーザー全件取得
func (t *User) List(c *gin.Context) {
	db := db.DB()

	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, users)
	}
}

// ユーザー作成
func (t *User) Create(c *gin.Context) {
	db := db.DB()

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// パスワードのハッシュ化
	hash_password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash_password)

	if err := db.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

// ユーザー更新
func (t *User) Update(c *gin.Context) {
	db := db.DB()

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	}
}

// func (t *User) Delete(c *gin.Context) {
// 	db := db.DB()

// 	var user model.User
// 	if err := db.Delete(&user, c.Param("id")).Error; err != nil {
// 		c.String(http.StatusInternalServerError, "Server Error")
// 	}
// }
