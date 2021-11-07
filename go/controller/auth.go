package controller

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

// envからトークンの鍵を取得
func getSecret() string {
	if err := godotenv.Load(".env/dev.env"); err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv("SECRET")
}

// ログイン
func (t *Auth) Login(c *gin.Context) {
	db := db.DB()
	var request_user model.User
	if err := c.BindJSON(&request_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// メールアドレスからユーザー取得
	var user model.User
	if err := db.First(&user, `email = ?`, request_user.Email).Error; err != nil {
		c.String(http.StatusInternalServerError, "user not found")
		return
	}
	// ユーザーパスワードの比較
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request_user.Password)); err != nil {
		c.String(http.StatusBadRequest, "Wrong password")
	} else {
		// トークンの作成
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid": user.ID,
			"iat": time.Now().Unix(),
		})
		secret := getSecret()
		str_token, _ := token.SignedString([]byte(secret))

		c.JSON(200, gin.H{
			"token": str_token,
		})
	}
}

// ログインした際にユーザー情報を送信
func (t *Auth) User(c *gin.Context) {
	// ヘッダーからID取得
	str_token := c.Request.Header["Authorization"]
	slice := strings.Split(str_token[0], " ")
	secret := getSecret()
	token, err := jwt.Parse(slice[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	user_claims := token.Claims
	user_map := user_claims.(jwt.MapClaims)
	user_id := user_map["uid"]

	// ユーザー情報を送信
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
