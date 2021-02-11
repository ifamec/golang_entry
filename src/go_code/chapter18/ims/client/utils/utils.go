package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"net"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("Client.process.readPkg : Wait For Server Message")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("Client.process.readPkg : Conn Read Msg Header Error -", err)
		// err = errors.New("read msg header error")
		return
	}
	// fmt.Println("Client.process : Read Success -", buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// read content
	n, err := conn.Read(buf[:pkgLen]) // conn.Read could be blocked only if both side is not closed
	if n != int(pkgLen) || err != nil {
		fmt.Println("Client.process.readPkg : Conn Read Msg Body Error -", err)
		// err = errors.New("read msg body error")
		return
	}
	// pkgLen to message.Message
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("Client.process.readPkg : Unmarshall Msg Error -", err)
		return
	}

	return
}
func writePkg(conn net.Conn, data []byte) (err error) {
	// 1. send length
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("Client : Conn Error -", err)
	}
	fmt.Println("Client : Msg Len Sent Success,", n, "Bytes Sent")

	// 7.2 send data
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("Client : Conn Error -", err)
	}
	// fmt.Println("Client : Msg Sent Success, Len", len(data), "Data:", string(data))
	fmt.Println("Client : Msg Sent Success")
	return
}
