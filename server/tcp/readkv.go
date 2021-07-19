package tcp

import (
	"bufio"
	"strconv"
)

func readLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}

	length, err := strconv.Atoi(tmp)
	if err != nil {
		return 0, err
	}
	return length, nil
}

func (s *Server) readKey(r *bufio.Reader) (string, error) {
	// key = bytes-array
	// bytes-array = <len><SP><content>
	klen, err := readLen(r)
	if err != nil {
		return "", err
	}

}

