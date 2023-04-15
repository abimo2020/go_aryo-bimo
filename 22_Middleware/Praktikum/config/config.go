package config

import (
	"fmt"
	"os"
	model "praktikum/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	config := Config{
		DB_Username: username,
		DB_Password: password,
		DB_Port:     port,
		// 192.168.0.107
		DB_Host: host,
		DB_Name: name,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{}, &model.Blog{}, &model.Book{})

	DB.Model(&model.Blog{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}
