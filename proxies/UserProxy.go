package proxies

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go-exercise/controllers"
	"go-exercise/vms"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserProxy struct {
	controllers.IUserController
}

func (proxy UserProxy) GetUserProfile(res http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	dto := proxy.GetUser(id)
	json.NewEncoder(res).Encode(vms.UserVM{dto.Id, dto.Name, dto.Email, dto.Avatar})
}

func (proxy UserProxy) UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var vm vms.UserUpdateVM
	json.Unmarshal(body, &vm)
	proxy.UpdateUser(id, vm.Name, vm.Email, vm.Avatar)
}
