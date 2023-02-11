package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	connectDB()
	db.AutoMigrate()
}

func connectDB() {
	var err error
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("PG_HOST"),
			os.Getenv("PG_PORT"),
			os.Getenv("PG_USER"),
			os.Getenv("PG_DB"),
			os.Getenv("PG_PASSWORD"),
		),
	)
	if err != nil {
		panic(err)
	}
}
