package main

import "fmt"

func InsertionSort(arr *[7]int) {
	for i := 1; i < len(arr); i++ {
		insertionVal := arr[i]
		insertionIndex := i - 1

		// lg -> sm
		for insertionIndex >= 0 && arr[insertionIndex] < insertionVal {
			arr[insertionIndex+1] = arr[insertionIndex]
			insertionIndex--
		}
		// insert
		if insertionIndex+1 != i {
			arr[insertionIndex+1] = insertionVal
		}
	}
}

func main() {
	var arr [7]int = [7]int{23, 0, 12, 56, 34, -1, 55}
	fmt.Println(arr)
	InsertionSort(&arr)
	fmt.Println(arr)
}
