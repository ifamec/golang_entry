package main

import "fmt"

type Cat struct {
	Num int
}

func (a Cat) test() { // Method belongs to type A
	a.Num = 2
	fmt.Println("a Cat.test: ", a.Num)
}

func main() {
	var a Cat = Cat{1}
	a.test()
	fmt.Println("main: ", a)
	// test() // Error
}
