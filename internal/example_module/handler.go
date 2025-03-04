package example_module

import "net/http"

type Handler interface {
	AddTicker() func(http.ResponseWriter, *http.Request)
}
