package main

import "fmt"

func main()  {
	num1 := 100
	fmt.Printf("%T, %v, %v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Printf("%T, %v, %v, %v\n", num2, num2, &num2, *num2) // *int, the address of *num2, address of num2, value of num2 pointing

}