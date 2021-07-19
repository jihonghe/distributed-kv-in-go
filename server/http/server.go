package http

import (
	cache2 "distributed-kv-in-go/server/cache"
	"net/http"
)

type Server struct {
	cache2.Cache
}

func (s *Server) Listen() {
	http.Handle("/api/cache/", s.cacheHandler())
	http.Handle("/api/status", s.statusHandler())
	if err := http.ListenAndServe(":12345", nil); err != nil {
		panic(err.Error())
	}
}

func New(c cache2.Cache) *Server {
	return &Server{c}
}
