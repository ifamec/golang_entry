package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	var key int
	var loop bool = true
	for loop {
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
			loop = false
		case 2:
			fmt.Println("Signup")
			loop = false
		case 3:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid, Retry")
		}
	}
	// show new info per user input
	if key == 1 {
		// login
		fmt.Printf("ID: ")
		fmt.Scanf("%d\n", &userId)
		fmt.Printf("Password: ")
		fmt.Scanf("%s\n", &userPwd)
		// login.go
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("Login Error:", err)
		}
	} else if key == 2 {
		fmt.Println("Sign Up")
	}
}
