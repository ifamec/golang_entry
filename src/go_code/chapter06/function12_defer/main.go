package main

import "fmt"

func sum(n1, n2 int) int {
	// when function is done, pop commands from defer stack
	defer fmt.Println("n1 =", n1) // 1. push 6. pop Print
	defer fmt.Println("n2 =", n2) // 2. push 5. pop Print
	n1++; n2++
	res := n1 + n2
	fmt.Println("n3 =", res)	// 3. push 4. pop Print
	return res
}

func main()  {
	fmt.Println("main:", sum(10, 20)) // 7. Print
}