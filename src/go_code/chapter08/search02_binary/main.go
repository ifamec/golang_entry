package main

import (
	"fmt"
)

func main() {
	var arr [7]int = [...]int{1, 3, 5, 7, 9, 11, 13}
	var target int = 13

	binarySearch(&arr, 0, len(arr) - 1, target)
}

func binarySearch (arr *[7]int, leftIndex int, rightIndex int, target int) {
	if leftIndex > rightIndex{
		fmt.Println("NOT FOUND")
		return
	}
	midIndex := (leftIndex + rightIndex) / 2

	if (*arr)[midIndex] > target {
		binarySearch(arr, 0, midIndex - 1, target)
	} else if (*arr)[midIndex] < target {
		binarySearch(arr, midIndex + 1, rightIndex, target)
	} else {
		fmt.Println("FOUNT, INDEX IS:", midIndex)
	}
}