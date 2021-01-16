package main

import "fmt"

func main()  {
	var p float32 = 5.12
	fmt.Println(p)

	var n1 float32 = -3.1E38
	var n2 float64 = -1.1E308
	fmt.Println(n1, n2)

	var n3 float32 = -10.00000301
	var n4 float64 = -10.00000301
	fmt.Println(n3, n4)

	var n5 = 1.1
	fmt.Printf("Data Type %T\n", n5)

	// decimal
	n6 := 5.12
	n7 := .123
	fmt.Println(n6, n7)

	// scientific notation
	n8 := 5.1234e2
	n9 := 5.1234E2
	n10 := 5.1234E-2
	fmt.Println(n8, n9, n10)
}
