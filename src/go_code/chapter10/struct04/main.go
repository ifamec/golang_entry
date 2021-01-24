package main

import "fmt"

type Cat struct {
	name string
	age int
	color string
	hobby string
}

func main() {
	var cat1 Cat = Cat{"Oolong", 4, "Orange", "Eat"}
	var cat2 *Cat = &cat1

	fmt.Println(cat1, cat2)
	fmt.Printf("cat1 add:%p, cat1 value:%v\ncat2 add%p, cat2 value%v\n", &cat1, cat1, &cat2, *cat2)
	(*cat2).name = "Latte"
	cat2.color = "Grey"
	fmt.Println(cat1, cat2)

}


