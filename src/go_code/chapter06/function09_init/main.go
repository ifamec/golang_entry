package main

import (
	"fmt"
	"go_code/chapter06/function09_init/utils"
)

var age int = test()

func test() int { // For checking global variable executes first
	fmt.Println("test(): 1")
	return 90
}

func main()  {
	fmt.Println("main(): 3")
	fmt.Println(utils.Age, utils.Name)
}

func init()  {
	fmt.Println("init(): 2")
}
