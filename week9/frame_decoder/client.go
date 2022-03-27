package week9

import (
	"encoding/binary"
	"fmt"
	"net"
)

func sendToServer(conn net.Conn) {
	// header : 64字节，代表内容长度
	message := "hello \n world\n"
	length := len(message)
	for {
		headerBuff := make([]byte, 64)
		binary.BigEndian.PutUint64(headerBuff, uint64(length))
		contentBuff := []byte(message)
		headerBuff = append(headerBuff, contentBuff...)
		_, err := conn.Write(headerBuff)
		if err != nil {
			fmt.Printf("write error %v\r\n", err)
		} else {
			fmt.Printf("write %s", contentBuff)
		}
	}
}
