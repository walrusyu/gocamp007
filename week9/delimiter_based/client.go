package week9

import (
	"fmt"
	"net"
)

func sendToServer(conn net.Conn) {
	message := "hello \n world\n"
	for {
		n, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("write error %v\r\n", err)
		} else {
			fmt.Printf("write %d bytes\r\n", n)
		}
	}
}
