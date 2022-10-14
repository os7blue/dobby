package api

var Apis = new(apiGroup)

type apiGroup struct {
	LoginApi loginApi
}
