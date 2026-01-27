package server

import "net/http"

type Server struct {
	port string
	router http.Handler
}

func New(port string, router http.Handler) *Server {
	return &Server{
		port: port,
		router: router,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(":" + s.port, s.router)
}
