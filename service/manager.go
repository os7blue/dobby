package service

var Services = new(serviceGroup)

type serviceGroup struct {
	AuthService authService
}
