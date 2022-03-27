package goim_decoder

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

// Package Length : 4bytes
// Header Length: 2bytes
// Protocol Version: 2bytes
// Operation: 4bytes
// Sequence Id: 4bytes
// Body: Package_Length - Header_Length

func decode(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		packageBuff := make([]byte, 4)
		reader.Read(packageBuff)
		packageLength := binary.BigEndian.Uint16(packageBuff)
		fmt.Printf("packageLength:%v\r\n", packageLength)

		headerBuff := make([]byte, 2)
		reader.Read(headerBuff)
		headerLength := binary.BigEndian.Uint16(headerBuff)
		fmt.Printf("headerLength:%v\r\n", headerLength)

		protocolBuff := make([]byte, 2)
		reader.Read(protocolBuff)
		protocolLength := binary.BigEndian.Uint16(protocolBuff)
		fmt.Printf("protocolLength:%v\r\n", protocolLength)

		operationBuff := make([]byte, 4)
		reader.Read(operationBuff)
		operationLength := binary.BigEndian.Uint16(operationBuff)
		fmt.Printf("operationLength:%v\r\n", operationLength)

		sequenceBuff := make([]byte, 4)
		reader.Read(sequenceBuff)
		sequenceLength := binary.BigEndian.Uint16(sequenceBuff)
		fmt.Printf("sequenceLength:%v\r\n", sequenceLength)

		bodyLength := packageLength - headerLength
		fmt.Printf("bodyLength:%v\r\n", bodyLength)
		bodyBuff := make([]byte, bodyLength)
		reader.Read(bodyBuff)
		fmt.Printf("read %s\r\n", bodyBuff)
	}
}

func encode(body string) []byte {
	headerLength := 16
	bodyLength := len(body)
	packageLength := headerLength + bodyLength
	packageBuff := make([]byte, packageLength)
	binary.BigEndian.PutUint16(packageBuff[:4], uint16(packageLength))
	binary.BigEndian.PutUint16(packageBuff[4:6], uint16(headerLength))
	binary.BigEndian.PutUint16(packageBuff[6:8], uint16(1))
	binary.BigEndian.PutUint16(packageBuff[8:12], uint16(1))
	binary.BigEndian.PutUint16(packageBuff[12:16], uint16(1))
	copy(packageBuff[16:], []byte(body))
	return packageBuff
}
