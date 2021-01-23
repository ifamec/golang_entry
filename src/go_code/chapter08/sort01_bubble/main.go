package main

import (
	"fmt"
)

func main() {
	var array [5]int = [5]int{3, 5, 1, 6, 9}
	bubbleSort(&array)
	fmt.Println(array)
}

func bubbleSort(arr *[5]int) {
	//for i := 0; i < len(*arr)-1; i++ {
	//	for j := i; j < len(*arr); j++ {
	//		if (*arr)[i] > (*arr)[j] {
	//			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	//		}
	//	}
	//}

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
