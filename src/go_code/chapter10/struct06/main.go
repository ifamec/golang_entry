package main

import "fmt"

type A struct {
	Num int
}
type B struct {
	Num int
}
type stu struct {
	name string
}
type student stu

func main() {
	var a A
	var b B
	a = A(b) // A and B have same fields
	fmt.Println(a, b)

	var stu0 stu
	var student0 student
	//stu = student // ERROR type different
	stu0 = stu(student0)
	fmt.Println(stu0)
}