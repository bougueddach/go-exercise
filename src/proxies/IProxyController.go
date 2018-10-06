package proxies

import "net/http"

type IProxyController interface {
	OK(res http.ResponseWriter, vm interface{})
	// add more to handle more cases if still time
}
