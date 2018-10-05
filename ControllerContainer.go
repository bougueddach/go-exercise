package main

import (
	"go-exercise/controllers"
	"go-exercise/proxies"
	"go-exercise/repositories"
)

type IServiceContainer interface {
	InjectUserProxy() proxies.UserProxy
}

type kernel struct{}

func (k *kernel) InjectUserProxy() proxies.UserProxy {

	userRepository := &repositories.UserRepository{}
	userController := &controllers.UserController{userRepository}
	userProxy := proxies.UserProxy{userController}

	return userProxy
}

func ControllerContainer() IServiceContainer {
	return &kernel{}
}
