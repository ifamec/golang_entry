package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var i int32 = 100

	// to float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("%v, %v, %v, %v\n", i, n1, n2, n3)

	fmt.Printf("%T, %T, %T, %T\n", i, n1, n2, n3)

	var m int64 = 10000
	var m1 int8 = int8(m)
	fmt.Printf("%v, %v\n", m, m1)
	fmt.Printf("%T, %T\n", m, m1)

	// data type to string
	var i0 int = 10
	var i1 float32 = 1.101
	var i2 bool = true
	var i3 byte = 'h'
	var str string

	str = fmt.Sprintf("%d", i0)
	fmt.Printf("type: %T, str: %q\n", str, str)

	str = fmt.Sprintf("%f", i1)
	fmt.Printf("type: %T, str: %q\n", str, str)

	str = fmt.Sprintf("%t", i2)
	fmt.Printf("type: %T, str: %q\n", str, str)

	str = fmt.Sprintf("%c", i3)
	fmt.Printf("type: %T, str: %q\n", str, str)

	fmt.Println("\n")

	// strconv
	str = strconv.FormatInt(int64(i0), 10)
	fmt.Printf("type: %T, str: %q\n", str, str)
	str = strconv.Itoa(i0)
	fmt.Printf("type: %T, str: %q\n", str, str)

	str = strconv.FormatFloat(float64(i1), 'f', 10, 64)
	fmt.Printf("type: %T, str: %q\n", str, str)

	str = strconv.FormatBool(i2)
	fmt.Printf("type: %T, str: %q\n", str, str)

	fmt.Println("\n")
	// string to basic data type

	var j0 string = "123450"
	var r0 int64
	r0, _ = strconv.ParseInt(j0, 10, 64)
	fmt.Printf("type: %T, str: %v\n", r0, r0)

	var j1 string = "10.123"
	var r1 float64
	r1, _ = strconv.ParseFloat(j1, 64)
	fmt.Printf("type: %T, str: %v\n", r1, r1)

	var j2 string = "false"
	var r2 bool
	r2, _ = strconv.ParseBool(j2)
	fmt.Printf("type: %T, str: %v\n", r2, r2)

	var j4 string = "Hello"
	var r4 int64
	var err error
	r0, err = strconv.ParseInt(j4, 10, 64)
	fmt.Printf("type: %T, str: %v\n", r4, r4)
	fmt.Println(err)
}