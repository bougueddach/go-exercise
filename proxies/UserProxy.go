package proxies

import (
	"encoding/json"
	"go-exercise/controllers"
	"go-exercise/vms"
	"io/ioutil"
	"net/http"
)

type UserProxy struct {
	controllers.IUserController
}

func (proxy UserProxy) GetUserProfile(res http.ResponseWriter, req *http.Request) {
	id := int64(1)
	dto := proxy.GetUser(id)
	json.NewEncoder(res).Encode(vms.UserVM{dto.Id, dto.Name, dto.Email, dto.Avatar})
}

func (proxy UserProxy) UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	id := int64(1)
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var vm vms.UserUpdateVM
	json.Unmarshal(b, &vm)
	proxy.UpdateUser(id, vm.Name, vm.Email, vm.Avatar)
}
