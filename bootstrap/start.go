package bootstrap

func Start() {

	//db connection init
	dbInit()
	//website and api init
	ginInit()

}
