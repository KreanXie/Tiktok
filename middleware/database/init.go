package database

import (
	"log"
	"tiktok/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:Krean751102!@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(
		&common.User{},
		&common.Video{},
		&common.UserVideoRelation{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
