package main

import "message-push/config"

func main() {

	config.DbInit()
	config.GinInit()

}
