package main

import (
	"encoding/binary"
	"encoding/json"
	_ "errors"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("Server.process.readPkg : Wait For Client Message")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("Server.process.readPkg : Conn Read Msg Header Error -", err)
		// err = errors.New("read msg header error")
		return
	}
	// fmt.Println("Client.process : Read Success -", buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// read content
	n, err := conn.Read(buf[:pkgLen]) // conn.Read could be blocked only if both side is not closed
	if n != int(pkgLen) || err != nil {
		fmt.Println("Server.process.readPkg : Conn Read Msg Body Error -", err)
		// err = errors.New("read msg body error")
		return
	}
	// pkgLen to message.Message
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("Server.process.readPkg : Unmarshall Msg Error -", err)
		return
	}

	return
}

// communication
func process(conn net.Conn) {
	// read client message
	defer conn.Close()

	// read client msg
	for {
		// pack read data as a function readPkg
		msg, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server.process : Client Exit")
				return
			} else {
				fmt.Println("Server.process : readPkg Error -", err)
				return
			}
		}
		fmt.Println("Server.process : readPkg Success -", msg)

	}
}

func main() {
	// listen port at 8889

	fmt.Println("Server : Start Listen At 8889")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("Server : Listen Error -", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("Server : Wait For Client")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Server : Listen Accept Error -", err)
		}

		// start goroutine if conn success
		go process(conn)
	}
}
