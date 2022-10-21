package router

var Routers = new(routerGroup)

type routerGroup struct {
	IndexRouter indexRouter
	AdminRouter adminRouter
}
