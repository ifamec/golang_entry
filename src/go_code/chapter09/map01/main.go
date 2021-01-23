package main

import "fmt"

func main() {
	var m1 map[string]string
	m1 = make(map[string]string, 10) // make to allocate ram
	m1["foo"] = "bar"
	m1["fo0"] = "bar"
	fmt.Println(m1)

	var m2 map[string]int = make(map[string]int)
	m2["Hey"] = 1
	m2["Hello"] = 2
	fmt.Println(m2)

	var m3 map[int]string = map[int]string{1: "Oh", 2: "My", 3: "God"}
	fmt.Println(m3)
	var m4 map[int]string = map[int]string{
		1: "Oh",
		2: "My",
		3: "Lord", // need "," here
	}
	fmt.Println(m4)

	var stu map[int]map[string]string = map[int]map[string]string{
		1: {"name": "Tom",   "Gender": "M"},
		2: {"name": "Jerry", "Gender": "M"},
		3: {"name": "Mary",  "Gender": "F"},
	}
	fmt.Println(stu)
}
