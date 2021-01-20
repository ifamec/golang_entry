package utils

import "fmt"

var Age int
var Name string
var Test int = utilTest()

func utilTest() int {
	fmt.Println("util var declare: 0.1")
	return 0
}

// need init in main.go then we could use

func init()  {
	Age = 10
	Name = "Jack"
	fmt.Println("util init(): 0.2")
}