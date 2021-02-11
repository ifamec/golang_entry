package main

import (
	"fmt"
	"go_code/chapter18/ims/client/processes"
	"os"
)

var userId int
var userPwd string

func main() {
	var key int
	for {
		fmt.Println()
		fmt.Println("-------- Welcome IMS --------")
		fmt.Println("\t 1. Login")
		fmt.Println("\t 2. Signup")
		fmt.Println("\t 3. Exit")
		fmt.Print("--------  Select: ")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("Login")
			// login
			fmt.Printf("ID: ")
			fmt.Scanf("%d\n", &userId)
			fmt.Printf("Password: ")
			fmt.Scanf("%s\n", &userPwd)
			// userProcess
			user := &processes.UserProcess{}
			err := user.Login(userId, userPwd)
			if err != nil {
				fmt.Println("Login Error:", err)
			}
		case 2:
			fmt.Println("Signup")
		case 3:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid, Retry")
		}
	}

}
