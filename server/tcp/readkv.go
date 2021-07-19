package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func readLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	tmp = strings.TrimSpace(tmp)

	length, err := strconv.Atoi(tmp)
	if err != nil {
		errInfo := fmt.Sprintf("failed to convert from str(value=%s) to int", tmp)
		log.Printf(errInfo)
		return 0, errors.New(errInfo)
	}
	return length, nil
}

// readKey
// @description 解析tcp-req-body中的ABNF表达式key
// key = bytes-array
// bytes-array = <len><SP><content>
func (s *Server) readKey(r *bufio.Reader) (string, error) {
	kLen, err := readLen(r)
	if err != nil {
		return "", err
	}

	key := make([]byte, kLen)
	//_, err = r.Read(key)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

// readKV
// @description 解析tcp-req-body中的ABNF表达式key-value
// key-value = <klen><SP><vlen><SP><kcontent><vcontent>
func (s *Server) readKV(r *bufio.Reader) (string, []byte, error) {
	kLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	vLen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}

	key := make([]byte, kLen)
	// bufio.Reader.Read(buf []byte): 读取的数据长度<=len(buf), 但要求读取的key是完整的，因此该方法不适用于读取key content
	//_, err = r.Read(key)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", nil, err
	}
	value := make([]byte, vLen)
	//_, err = r.Read(value)
	_, err = io.ReadFull(r, value)
	if err != nil {
		return "", nil, err
	}

	return string(key), value, nil
}
