package main

import "fmt"

func main()  {
	// basic data type in RAM
	var i int = 10
	// the ram address of i &i
	fmt.Println("value of i:", i)
	fmt.Println("ram address of i:", &i)

	// 1. ptr is an pointer variable
	// 2. ptr of ptr is *int
	// 3. the value of ptr is &i
	var ptr *int = &i
	fmt.Printf("ptr: %v \n", ptr)
	fmt.Printf("ptr address: %v \n", &ptr)
	fmt.Printf("ptr pointing value: %v\n", *ptr)

	fmt.Println("\nEXERCISE")
	exercise()
}

func exercise()  {
	var a int = 300
	fmt.Println("a value:", a)
	var ptr *int = &a
	fmt.Println("ptr:", ptr, "a address:",&a)
	*ptr = 400
	fmt.Println("val from *ptr:", *ptr, "a value after change:", a)
}