package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println(i)

	a := 9
	b := 6
	// switch a and b
	t := a
	a = b
	b = t
	fmt.Println(a, b)

	a += 7 // a = a + 7
	fmt.Println(a)


	// EXERCISE switch without temp
	fmt.Printf("\nEXERCISE\n")
	m := 5
	n := 6
	m = m + n
	n = m - n
	m = m - n
	fmt.Println(m, n)

	m, n = n, m
	fmt.Println(m, n)

}
