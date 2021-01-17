package main

import "fmt"

func main()  {
	var a int = 9
	var b int = 8

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)

	flag := a > b
	fmt.Println("flag =", flag)
}
