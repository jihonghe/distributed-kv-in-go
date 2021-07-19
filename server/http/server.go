package httpserver

import (
	"distributed-kv-in-go/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/api/cache/", s.cacheHandler())
	http.Handle("/api/status", s.statusHandler())
	if err := http.ListenAndServe(":12345", nil); err != nil {
		panic(err.Error())
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
