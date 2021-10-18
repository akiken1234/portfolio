package main

import (
	"github.com/akimotokensaku/portfolio/go/controller"
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	db.Connection()
	defer db.DB().Close()

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

	r.Run()
}
