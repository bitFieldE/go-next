package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	connectDB()
	Close()
}

func connectDB() {
	var err error
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
			"postgres",
			"password",
			"db",
			"5432",
			"go-next",
		),
	)
	if err != nil {
		panic(err)
	}
}

func Close() {
	defer db.Close()
}
