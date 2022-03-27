package week9

import (
	"fmt"
	"net"
)

func sendToServer(conn net.Conn) {
	for {
		buff := make([]byte, 1024)
		n, err := conn.Write(buff)
		if err != nil {
			fmt.Printf("write error %v\r\n", err)
		} else {
			fmt.Printf("write %d bytes\r\n", n)
		}
	}
}
