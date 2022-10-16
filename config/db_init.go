package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"message-push/bootstrap"
	"os"
)

var DB *gorm.DB

func DbInit() {
	db, err := gorm.Open(sqlite.Open(bootstrap.Option.DB.Url))
	if err != nil {
		fmt.Println("datasource init failed")
		os.Exit(1)
	}

	DB = db

}
