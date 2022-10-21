package bootstrap

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type email struct {
	Host     string   `ini:"host"`
	Port     string   `ini:"port"`
	Username string   `ini:"username"`
	Password string   `ini:"password"`
	Admin    []string `ini:"admin"`
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

var Option *option

func optionInit() {

	cfg, err := ini.Load("dev.ini")
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
