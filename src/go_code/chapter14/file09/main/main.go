package main

import (
	"flag"
	"fmt"
)

func main() {
	// define variables for get command params
	var user, pwd, host string
	var port int

	flag.StringVar(&user, "u", "", "User Name") // get user
	flag.StringVar(&pwd, "pwd", "", "Password") // get pwd
	flag.StringVar(&host, "h", "localhost", "Host") // get pwd
	flag.IntVar(&port, "port", 3306, "Port") // get pwd

	// transfer
	flag.Parse()

	fmt.Println("user: ", user)
	fmt.Println("pwd: ", pwd)
	fmt.Println("host: ", host)
	fmt.Println("port: ", port)
}
