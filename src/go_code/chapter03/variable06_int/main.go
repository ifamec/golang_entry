package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	// int8
	var i int8 = -128
	fmt.Println(i)

	// uint8
	var j uint8 = 255
	fmt.Println(j)

	var k int = 0x7FFFFFFFFFFFFFFF 	// int64
	var l uint = 0xFFFFFFFFFFFFFFFF // uint64
	var m rune = 0x7FFFFFF 			// int32
	var n byte = 0xFF 				// int32
	fmt.Println(k, l, m, n)

	var o = 100 // int
	fmt.Printf("Data Type of o is %T\n", o) // Format

	var o1 int64 = 100
	fmt.Printf("Data Type: %T  Byte Size: %d\n", o1, unsafe.Sizeof(o1))

	var age byte = 26
	fmt.Println("Age is:", age)



}
