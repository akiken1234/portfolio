package controller

import (
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

type JsonRequest struct {
	email    string `json:"email"`
	password string `json:"password"`
}

func (t *Auth) Login(c *gin.Context) {
	db := db.DB()

	var json JsonRequest
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(json)
	log.Println(json.email)
	log.Println(json.password)
	log.Println(c)
	log.Println(c.Param("req"))
	log.Println(c.Request)
	log.Println(c.Request.Context())
	log.Println(c.Request.Body)

	var user model.User
	err := db.First(&user, `email = ?`, "akimoto@gmail.com").Error

	if err != nil {
		c.String(http.StatusInternalServerError, "user not found")
		return
	}
	// DBから取得したユーザーパスワード(Hash)
	dbPassword := user.Password
	// フォームから取得したユーザーパスワード
	formPassword := "password4last"
	// formPassword := c.PostForm("password")
	log.Println(dbPassword)
	log.Println(formPassword)
	log.Println(c.Request.Body)
	// ユーザーパスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword))
	if err != nil {
		log.Println(111)
		c.String(http.StatusBadRequest, "Wrong password")
		c.Abort()
	} else {
		log.Println(222)
		// JWT
		claims := jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),            // stringに型変換
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := jwtToken.SignedString([]byte("secret"))
		if err != nil {
			log.Println(333)
			c.String(http.StatusInternalServerError, "Server Error")
		} else {
			log.Println(444)
			c.JSON(200, token)
		}
	}
}
