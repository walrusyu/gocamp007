package week9

import (
	"fmt"
	"net"
)

func readFromClient(conn net.Conn) {
	for {
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("read error %v\r\n", err)
		} else {
			fmt.Printf("read %d bytes\r\n", n)
		}
	}
}
