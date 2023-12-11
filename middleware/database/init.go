package database

import (
	"log"
	"tiktok/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn connect is hiden here
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(
		&common.User{},
		&common.Video{},
		&common.Like{},
		&common.View{},
		&common.Comment{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
