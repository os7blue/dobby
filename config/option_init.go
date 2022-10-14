package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type email struct {
	Host     string `ini:"host"`
	Username string `ini:"username"`
	Password string `ini:"password"`
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

type option struct {
	Email email
	DB    db
	Web   web
}

var Option option

func OptionInit() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("config file read failed")
		os.Exit(1)
	}

	Option := new(option)
	Option.Web.Port = "8081"
	Option.Web.Addr = "0.0.0.0"

	err = cfg.MapTo(Option)
	if err != nil {
		fmt.Println("config option init failed")
		os.Exit(1)
	}

}
