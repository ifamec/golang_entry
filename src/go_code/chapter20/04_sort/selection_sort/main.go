package main

import (
	"fmt"
)

func selectionSort(arr *[6]int) {
	for i := 0; i < len(arr) - 1; i++ {
		max := arr[i]
		maxIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
				maxIdx = j
			}
		}
		if maxIdx != i {
			arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
		}
	}
}

func main() {
	var arr [6]int = [6]int{10, 34, 19, 100, 80, 67}
	fmt.Println(arr)
	selectionSort(&arr)
	fmt.Println(arr)
}
