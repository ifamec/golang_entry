package main

import "fmt"

func main()  {
	var slice []int = make([]int, 4, 10)

	fmt.Println(slice)
	slice[1] = 1
	slice[2] = 2
	fmt.Println(slice, len(slice), cap(slice))

	var slice2 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice2, len(slice2), cap(slice2))
}