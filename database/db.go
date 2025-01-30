package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConfigDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/real_time_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection Error", err)
	}
	fmt.Println("Success Connect To Database")
	DB = db
}
