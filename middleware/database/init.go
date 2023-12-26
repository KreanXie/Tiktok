package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"tiktok/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

func InitDB() {
	var err error
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(
		&common.User{},
		&common.Video{},
		&common.View{},
		&common.Comment{},
		&common.Report{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() (DBConfig, error) {
	path := "config.json"
	var config DBConfig
	configFile, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)

	return config, err
}
