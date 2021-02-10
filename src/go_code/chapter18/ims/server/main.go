package main

import (
	"fmt"
	"net"
)

// communication
func process(conn net.Conn) {
	// read client message
	defer conn.Close()

	// read client msg
	for {
		buf := make([]byte, 4096)
		fmt.Println("Client.process : Wait For Client Message")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("Client.process : Read Error -", err)
			return
		}
		fmt.Println("Client.process : Read Success -", buf[:4])

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
