package proxies

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type ProxyController struct {
}

func (controller ProxyController) OK(res http.ResponseWriter, vm interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(vm)
}

func (controller ProxyController) ParseBody(r io.Reader, res http.ResponseWriter, structType interface{}) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		http.Error(res, "Error converting results to json", http.StatusInternalServerError)
	}
	err = json.Unmarshal(body, &structType)
	if err != nil {
		http.Error(res, "Bad request body", http.StatusBadRequest)
	}
}
