package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

type JsonRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const secret = "2FMd5FNSqS/nW2wWJy5S3ppjSHhUnLt8HuwBkTD6HqfPfBBDlykwLA=="

func (t *Auth) Login(c *gin.Context) {
	db := db.DB()

	var json JsonRequest
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	err := db.First(&user, `email = ?`, json.Email).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "user not found")
		return
	}
	// ユーザーパスワードの比較
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(json.Password))
	if err != nil {
		c.String(http.StatusBadRequest, "Wrong password")
	} else {
		// トークンの作成
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid": user.ID,
			"iat": time.Now().Unix(),
		})
		str_token, _ := token.SignedString([]byte(secret))

		c.JSON(200, gin.H{
			"token": str_token,
		})
	}
}

func (t *Auth) User(c *gin.Context) {
	// ヘッダーからID取得
	str_token := c.Request.Header["Authorization"]
	slice := strings.Split(str_token[0], " ")
	token, err := jwt.Parse(slice[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	user_claims := token.Claims
	user_map := user_claims.(jwt.MapClaims)
	user_id := user_map["uid"]

	// ユーザー送信
	db := db.DB()
	var user model.User
	err = db.First(&user, `id = ?`, user_id).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
	} else {
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}
