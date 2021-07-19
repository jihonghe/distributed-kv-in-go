package tcp

import (
	"fmt"
	"log"
	"net"
)

// sendResponse
// @description tcp回包：response = (bytes-array | error); error='-'<SP><content>
func sendResponse(value []byte, err error, conn net.Conn) error {
	// resp=-<elen><SP><econtent>
	if err != nil {
		errInfo := err.Error()
		resp := fmt.Sprintf("-%d %s", len(errInfo), errInfo)
		_, e := conn.Write([]byte(resp))
		return e
	}

	// resp=<vlen><SP><value>
	resp := fmt.Sprintf("%d %s", len(value), string(value))
	_, e := conn.Write([]byte(resp))
	log.Printf("send-resp: %s", resp)
	return e
}
