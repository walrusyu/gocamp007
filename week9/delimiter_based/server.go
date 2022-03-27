package week9

import (
	"bufio"
	"fmt"
	"net"
)

func readFromClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			fmt.Printf("read error %v\r\n", err)
		} else {
			fmt.Printf("read %s\r\n", line)
		}
	}
}
