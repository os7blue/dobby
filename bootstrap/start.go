package bootstrap

func Start() {

	//config option init ->  sort must first
	optionInit()
	//local cache init
	cacheInit()
	//db connection init
	dbInit()
	//auto create db table
	AutoTableInit()
	//website and api init
	ginInit()

}
