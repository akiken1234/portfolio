package main

import (
	"math/rand"
	"time"

	"github.com/akimotokensaku/portfolio/go/controller"
	"github.com/akimotokensaku/portfolio/go/db"
	"github.com/akimotokensaku/portfolio/go/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"os"
)

func router() {
	r := gin.Default()

	frontHost := "http://localhost:3000"
	// 本番環境の場合、envファイルからパスを取得
	if err := godotenv.Load(".env/dev.env"); err != nil {
		panic("Error loading .env file")
	}
	if production := os.Getenv("PRODUCTION"); production == "true" {
		frontHost = os.Getenv("FRONT_HOST")
	}

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
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			frontHost,
		},
	}))

	// ルーティング
	u := r.Group("/users")
	{
		ctrl := controller.NewUser()
		u.GET("", ctrl.List)
		u.POST("", ctrl.Create)
		u.PUT("", ctrl.Update)
		// u.DELETE("/:id", ctrl.Delete)
	}

	p := r.Group("/papers")
	{
		ctrl := controller.NewPaper()
		p.GET("", ctrl.List)
		p.POST("/get", ctrl.Get)
		p.POST("/create", ctrl.Create)
		p.POST("/download", ctrl.Download)
		p.DELETE("/:file_name", ctrl.Delete)
	}

	a := r.Group("/api/auth")
	{
		ctrl := controller.NewAuth()
		a.POST("/login", ctrl.Login)
		a.GET("/user", ctrl.User)
	}

	r.Run(":8080")
}

// マイグレーション
func migrate() {
	db.Connection()
	defer db.DB().Close()

	db.DB().AutoMigrate(&model.User{})
	db.DB().AutoMigrate(&model.Paper{})
}

// シーダー
func seed() {
	db := db.Connection()

	// userのseed
	var user [4]string
	user[0] = "曽倉 哲"
	user[1] = "真具 莉都"
	user[2] = "城 啓二"
	user[3] = "ナシーム・ニコラス・タレブ"

	count := 4
	for i := 0; i < count; i++ {
		name := user[i]
		email := fmt.Sprintf("email%d@gmail.com", i+1)
		password := fmt.Sprintf("password%d", i+1)
		hash_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		str_password := string(hash_password)
		if err != nil {
			return
		}
		user := model.User{Name: name, Email: email, Password: str_password}
		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}

	// paperのseed
	var title [15]string
	title[0] = "カント『純粋理性批判』における無限のアンチノミーについて"
	title[1] = "プラトンのイデア論について"
	title[2] = "アナクシマンドロスの「ト・アペイロン」について"
	title[3] = "カントールとゲーデルについて"
	title[4] = "カフカの『変身』について"
	title[5] = "ベケットの『ゴドーを待ちながら』について"
	title[6] = "ゴッホの『ひまわり』について"
	title[7] = "ライプニッツについて"
	title[8] = "ニーチェの超人概念について"
	title[9] = "ガタリの分裂分析におけるTFΦU図式について"
	title[10] = "ラマヌジャンの創造性について"
	title[11] = "血液型性格診断について"
	title[12] = "グロタンディークの晩年について"
	title[13] = "ボルヘスの文学について"
	title[14] = "後期のナボコフについて"

	var abstract [15]string
	abstract[0] = "カント『純粋理性批判』における第一アンチノミーである無限のアンチノミーを形而上学の視点から解釈する。"
	abstract[1] = "プラトンのイデア論は分析哲学からもポストモダン哲学からも批判されている。それらの批判がはたして妥当と言えるのか検証する。"
	abstract[2] = "初期ギリシア哲学者のアナクシマンドロスは、「万物はト・アペイロン（限りがないもの）」だと主張した。その可能性について分析する。"
	abstract[3] = "後年のカントールはシェイクスピア研究をした。また、後年のゲーデルは神の存在証明を試みた。その関係性について論じる。"
	abstract[4] = "カフカの『変身』を実存的視点から解釈する。"
	abstract[5] = "ベケットの『ゴドーを待ちながら』を時間論として解釈する。"
	abstract[6] = "ゴッホの『ひまわり』を超越的実体の表現として解釈する。"
	abstract[7] = "ライプニッツは、哲学、数学、神学などさまざまな分野で基本的な理論を発明した。その創造性について論じる。"
	abstract[8] = "ニーチェの超人概念について。超人はニーチェ本人からもニーチェ研究者からも抽象的に語られてきたが、具体例を示す。"
	abstract[9] = "ガタリの『分裂分析的地図作成法』におけるTFΦU図式について、形而上学的時間論として解釈する。"
	abstract[10] = "ラマヌジャンの数学的創造性と、ラマヌジャンが信仰していたヒンドゥー教の関連性について論じる。"
	abstract[11] = "血液型性格診断を科学的に検証する。"
	abstract[12] = "グロタンディークがなぜ何十年もピレネー山脈に隠棲したのかを考察する。"
	abstract[13] = "ボルヘスは多様な無限をオブジェのように表現した。それらの無限を分類してみる。"
	abstract[14] = "後期のナボコフの作品である『青白い炎』、『アーダ』、『透明な対象』を形而上学的に批評する。"

	paper := model.Paper{Title: title[0], Abstract: abstract[0], FileName: "file_name_1.pdf", UserId: 1}
	if err := db.Create(&paper).Error; err != nil {
		fmt.Printf("%+v", err)
	}
	count = 15
	for i := 1; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		user_id := (rand.Intn(4)) + 1
		file_name := fmt.Sprintf("file_name_%d.pdf", i+1)
		paper := model.Paper{Title: title[i], Abstract: abstract[i], FileName: file_name, UserId: user_id}
		if err := db.Create(&paper).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
}

func main() {
	migrate()
	seed()
	router()
}
