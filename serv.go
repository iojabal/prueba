package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func Newserver(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]
	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
