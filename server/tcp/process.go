package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (s *Server) process(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("close conn due to error: ", err)
			}
			return
		}
		switch op {
		default:
			log.Println("close conn due to invalid operation: ", err)
			return
		case 'S':
			err = s.set(conn, r)
		case 'G':
			err = s.get(conn, r)
		case 'D':
			err = s.del(conn, r)
		}
		if err != nil {
			log.Println("close conn due to error: ", err)
			return
		}
	}
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	key, value, err := s.readKV(r)
	if err != nil {
		return err
	}
	log.Printf("req-SET: key=%s, value=%s", key, string(value))
	if err = sendResponse(nil, s.Set(key, value), conn); err != nil {
		return err
	}
	return nil
}

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	value, err := s.Get(key)
	log.Printf("req-GET: key=%s, value=%s", key, value)
	if err = sendResponse(value, err, conn); err != nil {
		return err
	}
	return nil
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	key, err := s.readKey(r)
	if err != nil {
		return err
	}
	log.Printf("req-DEL: key=%s", key)
	if err = sendResponse(nil, s.Del(key), conn); err != nil {
		return err
	}
	return nil
}
