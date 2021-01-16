package main

import "fmt"

// Global Variable
var n0 = 100
var n00 = 200
var name0 = "Tim"
var (
	n000 = 300
	name00 = "Mary"
)

func main()  {
	//var n1 int
	//var n2 int
	//var n3 int
	var n1, n2, n3 int
	fmt.Println("n1, n2, n3", n1, n2, n3)

	// declare multiple var with different data type
	var n4, name, n5 = 100, "tom", 888
	fmt.Println("n4, name, n3", n4, name, n5)

	// use data type estimate
	n6, name2, n7 := 100, "jerry", 888
	fmt.Println("n6, name2, n7", n6, name2, n7)

	// Global Variable
	fmt.Println("Global Variable", n0, n00, name0, n000, name00)

}
