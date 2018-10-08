package proxies

import (
	"io"
	"net/http"
)

type IProxyController interface {
	OK(res http.ResponseWriter, vm interface{})
	ParseBody(r io.Reader, res http.ResponseWriter, structType interface{})
}
