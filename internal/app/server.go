package app

import "net/http"

type Server struct {
	mux    *http.ServeMux
	client *http.Server
}

// создаем новый сервер
func NewServer() *Server {
	mux := http.NewServeMux()

	return &Server{
		mux: mux,
		client: &http.Server{
			Addr:    ":3000",
			Handler: mux,
		},
	}
}

// запуск сервера
func (s *Server) Start() error {
	return s.client.ListenAndServe()
}

// добавления маршутов
func (s *Server) AddHandler(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(path, handler)
}
