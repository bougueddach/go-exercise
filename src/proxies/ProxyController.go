package proxies

import (
	"encoding/json"
	"net/http"
)

type ProxyController struct {
}

func (controller ProxyController) OK(res http.ResponseWriter, vm interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(vm)
}
