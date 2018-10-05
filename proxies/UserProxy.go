package proxies

import (
	"encoding/json"
	"go-exercise/controllers"
	"go-exercise/vms"
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
