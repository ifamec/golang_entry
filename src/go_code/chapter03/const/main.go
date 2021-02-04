package main

import "fmt"

func main() {
	var num int = 3
	const Pi = 3.14

	fmt.Println(num, Pi)

	// Pi = 10 // Error

	// const b = num / 3  // Error: num is not a const
	const b = 9 / 3
	fmt.Println(b)

	const (
		x = iota
		y
		z
	)
	fmt.Println(x, y, z)
}
