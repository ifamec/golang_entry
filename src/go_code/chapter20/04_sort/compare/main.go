package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(arr *[800000]int) {
	for i := 0; i < len(arr)-1; i++ {
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

func InsertionSort(arr *[800000]int) {
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

func QuickSort(left int, right int, arr *[800000]int) {
	l, r := left, right
	pivot := arr[(left+right)/2]
	for ; l < r; { // all smaller than pivot in left, others in right
		for ; arr[l] < pivot; {
			l++
		}
		for ; arr[r] > pivot; {
			r--
		}
		// if l >= r done in this round, else swap
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	var arr1 [800000]int
	var arr2 [800000]int
	var arr3 [800000]int
	var start, end int64
	for i := 0; i < 800000; i++ {
		number := rand.Intn(100000000)
		arr1[i] = number
		arr2[i] = number
		arr3[i] = number
	}
	start = time.Now().Unix()
	selectionSort(&arr1)
	end = time.Now().Unix()
	fmt.Println("Selection Sort Time:", end-start)

	start = time.Now().Unix()
	InsertionSort(&arr2)
	end = time.Now().Unix()
	fmt.Println("Insertion Sort Time:", end-start)

	start = time.Now().Unix()
	QuickSort(0, len(arr3)-1, &arr3)
	end = time.Now().Unix()
	fmt.Println("Quick Sort Time:", end-start)
}
