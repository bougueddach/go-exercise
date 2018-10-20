package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

type IRouter interface {
	InitRouter() http.Handler
}

type router struct{}

func (router *router) InitRouter() http.Handler {
	r := chi.NewRouter()
	router.initUserProxyRoutes(r)
	return r
}

func (router *router) initUserProxyRoutes(r *chi.Mux) {
	userProxy := ControllerContainer().InjectUserProxy()
	r.Get("/user/{id}", userProxy.GetUserProfile)
	r.Post("/user/{id}", userProxy.UpdateUserProfile)
}

func Router() IRouter {
	return &router{}
}
