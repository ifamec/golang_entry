package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Client : Conn Error", err)
		return
	}
	defer conn.Close()
	fmt.Println("Client : Conn Success", conn, conn.LocalAddr().String())

	for {
		// 1. Get input and send to server
		reader := bufio.NewReader(os.Stdin) // terminal input

		line, err := reader.ReadString('\n') // Read user input
		if err != nil {
			fmt.Println("Client : Read String Error", err)
		}

		line = strings.Trim(line, " \n\r")
		if line == "exit" {
			fmt.Println("Client : Exit")
			break
		}
		// send data
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("Client : Conn Write Error", err)
		}
		fmt.Println("Client : Conn Write Success", n, "Bytes Sent.")

		buf := make([]byte, 1024)
		n, err = conn.Read(buf) // block here if no write on client side
		if err == io.EOF {
			fmt.Println("Client : Read Error")
			return
		}
		fmt.Print("Server Reply: ", string(buf[:n])) // !!!! data length

	}
}
