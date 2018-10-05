package main

import (
	"go-exercise/controllers"
	"go-exercise/proxies"
	"go-exercise/repositories"
	"net/http"
	"sync"
)

func main()  {
	userProxy := ServiceContainer().InjectUserProxy()
	http.HandleFunc("/user", userProxy.GetUserProfile)
	host := "localhost:8080"
	runServer(host)
}

func runServer(host string) error {
	return http.ListenAndServe(host, nil)
}

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

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}