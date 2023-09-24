package util

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbName := os.Getenv("DB_FILE")
	if len(dbName) <= 0 {
		log.Fatal("db name not set.")
	}

	DB, err = gorm.Open(sqlite.Open(dbName))
	if err != nil {
		log.Fatal("failed to connect to db")
	}
}
