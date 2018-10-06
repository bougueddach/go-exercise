package main

import (
	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	r := chi.NewRouter()
	router.InitUserProxyRoutes(r)
	return r
}

func (router *router) InitUserProxyRoutes(r *chi.Mux) {
	userProxy := ControllerContainer().InjectUserProxy()
	r.Get("/user/{id}", userProxy.GetUserProfile)
	r.Post("/user/{id}", userProxy.UpdateUserProfile)
}

func Router() IChiRouter {
	return &router{}
}
