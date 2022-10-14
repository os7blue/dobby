package main

import "message-push/config"

func main() {

	//config option init ->  sort must first
	defer config.OptionInit()
	//db connection init
	defer config.DbInit()
	//website and api init
	defer config.GinInit()

}
