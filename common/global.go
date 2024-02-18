package common

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

type email struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Admin    string `ini:"admin"`
}

type db struct {
	Url      string `ini:"url"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type web struct {
	Port string `ini:"port"`
	Addr string `ini:"addr"`
}

type setting struct {
	PushTitle string `ini:"pushTitle"`
}

type option struct {
	Email   email
	DB      db
	Web     web
	Setting setting
}

var Option *option

func init() {

	cfg, err := ini.Load("/Users/codecv/data/code/workplace/go/dobby/dev.ini")
	if err != nil {
		fmt.Println("config file load failed")
		os.Exit(1)
	}

	op := new(option)
	op.Web.Addr = "0.0.0.0"
	op.Web.Port = "8081"
	err = cfg.MapTo(op)
	if err != nil {
		fmt.Println("config init failed")
		os.Exit(1)
	}

	Option = op

}

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open(Option.DB.Url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
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

var LocalCache *ristretto.Cache

func init() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		fmt.Println("local cache init failed")
		os.Exit(1)
	}

	LocalCache = cache

}
