package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var database *gorm.DB

func Connection() *gorm.DB {
	dbDriver := "mysql"
	dbUser := "user"
	dbPass := "password"
	dbAddress := "db"
	dbName := "portfolio"

	// 本番環境の場合、envファイルからパスを取得
	if err := godotenv.Load(".env/dev.env"); err != nil {
		panic("Error loading .env file")
	}
	if production := os.Getenv("PRODUCTION"); production == "true" {
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbAddress = os.Getenv("DB_ADDRESS")
	}

	var err error
	database, err = gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbAddress+":3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	database.LogMode(true)
	return database
}

func DB() *gorm.DB {
	return database
}
