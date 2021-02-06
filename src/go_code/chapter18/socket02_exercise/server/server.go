package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func Process(conn net.Conn)  {
	defer conn.Close()

	for { // get data from client
		buf := make([]byte, 1024)
		// fmt.Println("Server : Waiting", conn.RemoteAddr().String(), "Input...")
		n, err := conn.Read(buf) // block here if no write on client side
		if err == io.EOF {
			fmt.Println("Server : Client", conn.RemoteAddr().String(), "Close.")
			return
		}
		// address := conn.RemoteAddr()
		var str string = string(buf[:n])
		str = strings.Trim(str, " \n\r")

		var rtnval string
		// fmt.Println("Server : Accept New Conn", conn, conn.RemoteAddr().String())
		switch str {
		case "Hello": 	rtnval = "Hello, how are you."
		case "Weather": rtnval = "Nice weather today."
		default: 		rtnval = "Sorry I dont understand."
		}
		// if err != nil {
		// 	fmt.Println("Server : Conn Read Error", err, conn.RemoteAddr().String())
		// 	return
		// }
		// show client input
		fmt.Println(conn.RemoteAddr().String(), " Input : ", str, " Reply: ", rtnval) // !!!! data length
		// fmt.Println("Server : Conn Read Success - ", n, "Bytes Received.")
		// send back
		_, err = conn.Write([]byte(rtnval + "\n"))
		if err != nil {
			fmt.Println("Server : Reply Error", err)
		}
	}
}

func main() {
	fmt.Println("Server : Setup Listener")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Server : Listen Error", err)
		return
	}
	defer listen.Close()

	// Accept() wait for new connection
	for {
		fmt.Println("Server : Wait For Client Connection...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Server : Accept Error", err)
		} else {
			fmt.Println("Server : Accept New Conn", conn, conn.RemoteAddr().String())
		}

		// Prepare goroutine to serve client
		go Process(conn)
	}

	// fmt.Println("Listen", listen)

}
