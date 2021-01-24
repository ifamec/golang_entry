package main

import "fmt"

type Cat struct {
	name string
	age int
	color string
	hobby string
}

func main() {
	var cat1 Cat
	fmt.Printf("%p\n", &cat1)
	cat1 = Cat{"Oolong", 4, "Orange", "Sleep"}
	var cat2 Cat = Cat{"Latte", 4, "Grey", "Eat"}
	fmt.Println(cat1, cat2)
}
