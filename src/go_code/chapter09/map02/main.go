package main

import "fmt"

func main() {
	var m map[int]string = map[int]string{
		1: "Oh",
		2: "My",
		3: "Lord", // need "," here
	}
	fmt.Println(m)
	// add
	m[4] = "!!!"
	// modify
	m[3] = "God"
	fmt.Println(m)
	// delete
	delete(m, 4)
	delete(m, 5)
	fmt.Println(m)
	// search
	value, find := m[3]
	fmt.Println(value, find)
	value, find = m[4]
	fmt.Println(value, find) // use default in value if key not exist

	// delete all
	m = make(map[int]string)
	fmt.Println(m)
}
