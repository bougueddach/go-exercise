package proxies

import (
	"encoding/json"
	"go-exercise/controllers"
	"net/http"
)

type UserProxy struct {
	controllers.IUserController
}

func (proxy UserProxy) GetUserProfile(res http.ResponseWriter, req *http.Request) {
	id := int64(0)
	user := proxy.GetUser(id)
	json.NewEncoder(res).Encode(user)
}
