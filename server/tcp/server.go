package tcp

import (
	"distributed-kv-in-go/server/cache"
	"log"
	"net"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", ":12346")
	if err != nil {
		panic(err)
	}

	for {
		log.Println("ready to accept a connection")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("accept a connection, ready to process")
		go s.process(conn)
	}
}
