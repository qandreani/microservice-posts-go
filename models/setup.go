package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

var DB *gorm.DB

const (
	host     = "51.159.11.187"
	port     = 30902
	user     = "quentin"
	password = "@Quentin23"
	dbname   = "posts"
)

func ConnectDatabase() {
	prosgret_conname := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Post{})

	DB = database
}