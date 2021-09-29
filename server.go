package main

import (
	"github.com/gorilla/mux"
)


type Server struct {
	*mux.Router
	Documents []Doc
}

// NewServer initialising server
func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Documents: []Doc{},
	}
	s.requests()
	return s
}
