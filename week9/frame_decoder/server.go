package week9

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

func readFromClient(conn net.Conn) {
	// header : 64字节，代表内容长度
	reader := bufio.NewReader(conn)
	for {
		headerBuff := make([]byte, 64)
		_, err := reader.Read(headerBuff)
		if err != nil {
			fmt.Printf("read header error %v\r\n", err)
			return
		}
		contentLength := binary.BigEndian.Uint64(headerBuff)
		var contentBuff = make([]byte, contentLength)
		_, err = reader.Read(contentBuff)
		if err != nil {
			fmt.Printf("read content error %v\r\n", err)
		} else {
			fmt.Printf("read %s\r\n", contentBuff)
		}
	}
}
