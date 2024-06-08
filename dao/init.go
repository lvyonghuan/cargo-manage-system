package dao

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(username, password string) {
	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/cargo_management"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		time.Sleep(10 * time.Second)
		panic(err)
	}
}
