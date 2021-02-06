package main

import "fmt"

func login(userId int, userPwd string) (err error) {
	fmt.Println("Your Input:", userId, userPwd)
	return nil
}
