package main

import "fmt"

func main() {
	var m map[int]string = map[int]string{
		1: "Oh",
		2: "My",
		3: "Lord", // need "," here
	}
	for key, value := range m {
		fmt.Printf("KEY: %v\t VALUE: %v\n", key, value)
	}
	fmt.Println(len(m))

	var stu map[int]map[string]string = map[int]map[string]string{
		1: {"name": "Tom",   "Gender": "M"},
		2: {"name": "Jerry", "Gender": "M"},
		3: {"name": "Mary",  "Gender": "F"},
	}
	for key, value := range stu {
		fmt.Printf("KEY: %v\n", key)
		for k, v := range value {
			fmt.Printf("\tKEY: %v\tVALUE: %v\n", k, v)
		}
	}
}
