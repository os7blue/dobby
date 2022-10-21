package bootstrap

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func dbInit() {
	db, err := gorm.Open(sqlite.Open(Option.DB.Url))
	if err != nil {
		fmt.Println("datasource init failed")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("datasource init failed")
		os.Exit(1)
	}

	DB = db

}
