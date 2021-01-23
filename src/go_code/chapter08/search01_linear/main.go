package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var arr [4]string = [4]string{"AA", "BB", "CC", "DD"}

	var target string
	fmt.Println("INPUT STRING:")
	fmt.Scanf("%s", &target)

	fmt.Println("INPUT IS", target, "\t", func() (str string) {
		index := linearSearch(&arr, target)
		if index == -1 {
			return "NOT FOUND IN ARRAY"
		}
		return "FOUND, INDEX IS " + strconv.Itoa(index)
	}())
}

func linearSearch(arr *[4]string, target string) (idx int) {
	idx = -1
	for index, value := range *arr {
		if value == target {
			idx = index
		}
	}
	return
}