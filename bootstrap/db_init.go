package bootstrap

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"message-push/common"
	"message-push/model"
	"os"
)

var DB *gorm.DB

func dbInit() {
	db, err := gorm.Open(sqlite.Open(common.Option.DB.Url))
	if err != nil {
		fmt.Println("datasource init failed")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("datasource init failed")
		os.Exit(1)
	}
	autoCreateTable(db)
	DB = db

}

func autoCreateTable(db *gorm.DB) {

	err := db.AutoMigrate(
		&model.PushChannelInfo{},
		&model.PushChannelLine{},
	)
	if err != nil {
		return
	}

}
