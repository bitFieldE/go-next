package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bitFieldE/go-next/backend/config"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var err error
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
			config.Db.User,
			config.Db.Password,
			config.Db.Host,
			config.Db.Port,
			config.Db.Name,
		),
	)
	if err != nil {
		log.Printf("%s: could open PostgreSQL server", err)
		panic(err)
	}

	return db
}
