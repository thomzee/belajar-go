package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB
var err error

func InitDb() {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		port     = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_NAME")
	)

	dsn := user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection Error:", err.Error())
		return
	}
}

func GetDb() *gorm.DB {
	return db
}
