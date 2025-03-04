package delivery

import (
	"net/http"
	"template/internal/example_module"
)

func MapRoutes(addHandlerFunc func(path string, handler func(http.ResponseWriter, *http.Request)), h example_module.Handler) {
	addHandlerFunc("/", h.AddTicker())
}
