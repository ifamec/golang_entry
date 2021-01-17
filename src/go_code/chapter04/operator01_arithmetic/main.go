package main

import "fmt"

func main() {
	// /
	fmt.Println(10 / 4) // if both of the number is int -> keep int part only
	var n1 float32 = 10 / 4
	fmt.Println(n1) // Still 2

	// to keep floating result:
	var n2 float32 = 10. / 4
	fmt.Println(n2)

	// %
	//a % b = a - a / b * b
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) // -1

	var i int = 10
	i++ // i = i + 1
	fmt.Println(i)
	i-- // i = i - 1
	fmt.Println(i)

	//j := i++ // ERROR
	i++
	j := i
	fmt.Println(j)

	//if i++ > 0 {
	//
	//} // ERROR

	//++i // ERROR
	exercise()
}

func exercise() {
	fmt.Printf("\nEXERCISE\n")
	// Calculate days to weeks and days
	var total int = 97
	fmt.Printf("97 days = %v weeks and %v days.\n", total / 7, total % 7)

	// Fahrenheit to Celsius
	var f float32 = 72
	fmt.Printf("%v Fahrenheit = %v Celsius,\n", f, 5. / 9 * (f - 32))
}
