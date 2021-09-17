package main

import (
	"local.packages/controller"
	"local.packages/db"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db.Connection()
	defer db.DB().Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hoge",
		})
	})

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
