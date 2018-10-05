package main

import (
	"github.com/go-chi/chi"
	"go-exercise/controllers"
	"go-exercise/proxies"
	"go-exercise/repositories"
	"net/http"
	"sync"
)

func main() {
	userProxy := ServiceContainer().InjectUserProxy()
	r := chi.NewRouter()
	r.Get("/user/id", userProxy.GetUserProfile)
	r.Post("/user/id", userProxy.UpdateUserProfile)
	host := "localhost:8080"
	http.ListenAndServe(host, r)
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
