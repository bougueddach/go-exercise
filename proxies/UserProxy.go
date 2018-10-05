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
		http.Error(res, "Bad url param", http.StatusBadRequest)
	}
	dto := proxy.GetUser(id)
	json.NewEncoder(res).Encode(vms.UserVM{dto.Id, dto.Name, dto.Email, dto.Avatar})
}

func (proxy UserProxy) UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, "Bad url param", http.StatusBadRequest)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error converting results to json", http.StatusInternalServerError)
	}

	var vm vms.UserUpdateVM
	err = json.Unmarshal(body, &vm)
	if err != nil {
		http.Error(res, "Bad request body", http.StatusBadRequest)
	}
	proxy.UpdateUser(id, vm.Name, vm.Email, vm.Avatar)
}
