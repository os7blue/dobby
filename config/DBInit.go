package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

type dbConfig struct {
	url      string `ini:"url"`
	username string `ini:"username"`
	password string `ini:"password"`
}

func DbInit() *gorm.DB {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("config file read failed")
		os.Exit(1)
	}

	dbc := new(dbConfig)
	err = cfg.Section("Db").MapTo(dbc)
	if err != nil {
		fmt.Println("db config read failed")
		os.Exit(1)
	}

	DB, err := gorm.Open(sqlite.Open(dbc.url))
	if err != nil {
		fmt.Println("datasource init failed")
		os.Exit(1)
	}

	return DB

}
