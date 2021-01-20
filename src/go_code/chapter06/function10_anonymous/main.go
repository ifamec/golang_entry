package main

import "fmt"

var (
	Fn1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main()  {
	// anonymous func - use at declaration
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	// assign anonymous function to a variable
	diff := func(n1 int, n2 int) int {
		return n1 - n2
	}
	// data tyoe func  func(int, int) int
	res2 := diff(10, 30)
	fmt.Printf("%T, %d\n", diff, res2)

	// global anonymous function
	res3 := Fn1(4, 9)
	fmt.Println(res3)
}