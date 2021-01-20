package main

import (
	"fmt"
	"strings"
)

func accumulator() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}
func accumulator2() func(int) int {
	var n int = 10
	var str string = "Hello"
	return func(x int) int {
		n += x
		str += string(rune(36))
		fmt.Println("str =", str) // str="Hello$" str="Hello$$" str="Hello$$$"
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if ! strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func main()  {

	f := accumulator()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
	f2 := accumulator2()
	fmt.Println(f2(1)) // 11
	fmt.Println(f2(2)) // 13
	fmt.Println(f2(3)) // 16

	closure := makeSuffix(".jpg")
	fmt.Println(closure("aaa"))
	fmt.Println(closure("aaa.jpg"))
}