package delivery

import (
	"net/http"
	"template/internal/example_module"
)

type Handler struct {
	uc example_module.UseCase
}

func NewHandler(uc example_module.UseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) AddTicker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ticker := r.Context().Value("ticker").(string)
		resp, err := h.uc.ProcessTicker(ticker)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	}
}
