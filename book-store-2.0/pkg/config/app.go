package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "my_db"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "username:@tcp(127.0.0.1:3306)/book-store-2.0?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
