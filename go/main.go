package main

import (
	"github.com/akimotokensaku/portfolio/go/controller"
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"

	"fmt"
)

// ルーティング
func router() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			// "Access-Control-Allow-Headers", "*",
			"Access-Control-Allow-Origin", // 追加
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://host.docker.internal:3000",
		},
	}))

	u := r.Group("/users")
	{
		ctrl := controller.NewUser()
		u.GET("", ctrl.List)
		u.GET("/:id", ctrl.Get)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	p := r.Group("/papers")
	{
		ctrl := controller.NewPaper()
		p.GET("", ctrl.List)
		p.POST("/get", ctrl.Get)
		p.POST("", ctrl.Create)
		p.DELETE("/:id", ctrl.Delete)
	}

	c := r.Group("/comments")
	{
		ctrl := controller.NewComment()
		c.GET("", ctrl.List)
		c.GET("/:id", ctrl.Get)
		c.POST("", ctrl.Create)
		c.DELETE("/:id", ctrl.Delete)
	}

	m := r.Group("/messages")
	{
		ctrl := controller.NewMessage()
		m.GET("", ctrl.List)
		m.GET("/:id", ctrl.Get)
		m.POST("", ctrl.Create)
		m.DELETE("/:id", ctrl.Delete)
	}

	a := r.Group("/api/auth/")
	{
		ctrl := controller.NewAuth()
		a.POST("/login", ctrl.Login)
		a.GET("/user", ctrl.User)
	}

	r.Run()
}

// マイグレーション
func migrate() {
	db.Connection()
	defer db.DB().Close()

	db.DB().AutoMigrate(&model.User{})
	db.DB().AutoMigrate(&model.Paper{})
	db.DB().AutoMigrate(&model.Message{})
	db.DB().AutoMigrate(&model.Comment{})
}

// シーダー
func seed() {
	db := db.Connection()

	// userのseed
	count := 3
	for i := 1; i <= count; i++ {
		name := fmt.Sprintf("name%d", i)
		email := fmt.Sprintf("email%d@gmail.com", i)
		password := fmt.Sprintf("password%d", i)
		hash_password, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

		user := model.User{Name: name, Email: email, Password: hash_password}
		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}

	// paperのseed
	papers1 := model.Paper{Title: "アリストテレスについて", Abstract: "概要1", FileName: "file_name_1.pdf", UserId: 1}

	if err := db.Create(&papers1).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	papers2 := model.Paper{Title: "カントについて", Abstract: "概要2", FileName: "file_name_2.pdf", UserId: 2}

	if err := db.Create(&papers2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	papers3 := model.Paper{Title: "ニーチェについて", Abstract: "概要3", FileName: "file_name_3.pdf", UserId: 3}

	if err := db.Create(&papers3).Error; err != nil {
		fmt.Printf("%+v", err)
	}
}

func main() {

	migrate()
	seed()
	router()

}
