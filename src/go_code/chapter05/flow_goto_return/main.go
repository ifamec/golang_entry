package main

import "fmt"

func main()  {
	fmt.Println("1")
	goto lable1
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	lable1: fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	return
	fmt.Println("8")
}