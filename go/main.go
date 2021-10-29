package main

import (
	"github.com/akimotokensaku/portfolio/go/controller"
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
	"time"
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
		MaxAge: 24 * time.Hour,
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
		p.GET("/:id", ctrl.Get)
		p.POST("", ctrl.Create)
		p.POST("/download", ctrl.Download)
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
		// a.GET("/user", ctrl.user)
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
	users1 := model.User{Name: "akimoto", Email: "akimoto@gmail.com", Password: "password1", AdminRole: true}

	if err := db.Create(&users1).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	users2 := model.User{Name: "suzuki", Email: "suzuki@gmail.com", Password: "password2"}

	if err := db.Create(&users2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	users3 := model.User{Name: "satou", Email: "satou@gmail.com", Password: "password3"}

	if err := db.Create(&users3).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	// paperのseed
	papers1 := model.Paper{Title: "アリストテレスについて", Abstract: "概要１", FileName: "file_name_1.pdf", UserId: 1}

	if err := db.Create(&papers1).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	papers2 := model.Paper{Title: "カントについて", Abstract: "概要２", FileName: "file_name_2.pdf", UserId: 2}

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
