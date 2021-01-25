package main

import "fmt"

type Point struct {
	x int
	y int
}

func main()  {
	var a interface{}
	var point Point = Point{1, 2}
	a = point // √
	var b Point
	// b = a // Error a is not a Point need type assertion
	b = a.(Point)
	fmt.Println(b)


	// var x interface{}
	// var b2 float64 = 1.1
	// x = b2
	// var c float64 = x.(float64)
	// fmt.Printf("%T, %v\n", c, c) // float64 1.1

	// type assertion with check
	var x interface{}
	var b2 float64 = 1.1
	x = b2
	c, flag := x.(float64)
	if flag {
		fmt.Printf("Assertion Success: %T, %v\n", c, c)
	} else {
		fmt.Println("Assertion Fails")
	}

	if d, flag := x.(float32); flag {
		fmt.Printf("Assertion Success: %T, %v\n", d, d)
	} else {
		fmt.Println("Assertion Fails")
	}
	fmt.Println("Fin")
}