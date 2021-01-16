package main

import "fmt"

func main()  {
	var i int  = 10
	i = 30
	i = 50

	fmt.Println(i)

	//i = 1.2 // Error - constant 1.2 truncated to integer

	//var i int = 100 	// redeclared
	//i := 100 			// redeclared
}
