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
	cat1 = Cat{}
	cat1.name = "Oolong"
	cat1.age = 4
	cat1.color = "Orange"
	cat1.hobby = "Sleep"

	var cat2 Cat = Cat{"Latte", 4, "Grey", "Eat"}

	fmt.Println(cat1, cat2)

	var cat3 *Cat = new(Cat)
	(*cat3).name = "C3"
	cat3.age = 5 // Simplified
	(*cat3).color = "Ame"
	cat3.hobby = "Climb" // Simplified
	fmt.Println(cat3, *cat3)

	var cat4 *Cat = &Cat{}
	(*cat4).name = "C4"
	cat4.age = 5 // Simplified
	fmt.Println(*cat4)
}


