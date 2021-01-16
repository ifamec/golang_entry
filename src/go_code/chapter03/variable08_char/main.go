package main

import "fmt"

func main()  {
	var c1 byte = 'a'
	var c2 byte = '0'

	// byte value -> ascii order
	fmt.Println(c1, c2)
	// print char
	fmt.Printf("%c %c\n", c1, c2)

	//var c3 byte = '北' // Error Overflow
	var c3 rune = '北'
	fmt.Printf("%c, %d\n", c3, c3)

	var c4 int = 22269
	var c5 byte = 120
	fmt.Printf("%c, %d, %c, %d\n", c4, c4, c5 ,c5)

	var n1 = 10 + 'a'
	fmt.Printf("%c, %d", n1, n1)


}
