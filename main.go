package main

import (
	"message-push/bootstrap"
	"message-push/config"
)

func main() {

	//config option init ->  sort must first
	bootstrap.OptionInit()
	//local cache init
	bootstrap.CacheInit()
	//db connection init
	config.DbInit()
	//website and api init
	config.GinInit()

}
