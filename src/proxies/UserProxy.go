package proxies

import (
	"github.com/go-chi/chi"
	"go-exercise/src/controllers"
	"go-exercise/src/vms"
	"net/http"
	"strconv"
)

type UserProxy struct {
	controllers.IUserController
	IProxyController
}

func (proxy UserProxy) GetUserProfile(res http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, "Bad url param", http.StatusBadRequest)
	}
	dto := proxy.GetUser(id)
	vm := vms.UserVM{dto.Id, dto.Name, dto.Email, dto.Avatar}
	proxy.OK(res, vm)
}

func (proxy UserProxy) UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, "Bad url param", http.StatusBadRequest)
	}
	vm := vms.UserUpdateVM{}
	proxy.ParseBody(req.Body, res, &vm)
	proxy.UpdateUser(id, vm.Name, vm.Email, vm.Avatar)
}
