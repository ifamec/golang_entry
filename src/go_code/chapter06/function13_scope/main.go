package main

import "fmt"

var age int = 22		// scope: current package
var Name string = "Tom" // scope: project

//Score := 10 // Error

// local variable
func local()  {
	var age int = 10
	var Name string = "Jack"
	fmt.Println(age, Name)
}

func main()  {

	fmt.Println(age, Name)
	local()
}