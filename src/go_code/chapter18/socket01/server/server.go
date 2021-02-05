package main

import (
	"fmt"
	"io"
	"net"
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
		// if err != nil {
		// 	fmt.Println("Server : Conn Read Error", err, conn.RemoteAddr().String())
		// 	return
		// }
		// show client input
		fmt.Print(conn.RemoteAddr().String(), " Input : ", string(buf[:n])) // !!!! data length
		// fmt.Println("Server : Conn Read Success - ", n, "Bytes Received.")
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
