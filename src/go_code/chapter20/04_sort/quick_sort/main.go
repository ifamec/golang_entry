package main

import "fmt"

func QuickSort(left int, right int, arr *[6]int) {
	l, r := left, right
	pivot := arr[(left + right) / 2]
	for ; l < r; { // all smaller than pivot in left, others in right
		for ; arr[l] < pivot; { l++ }
		for ; arr[r] > pivot; { r-- }
		// if l >= r done in this round, else swap
		if l >= r { break }
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot { r-- }
		if arr[r] == pivot { l++ }
	}
	if l == r { l++; r-- }
	if left < r  { QuickSort(left, r, arr)  }
	if right > l { QuickSort(l, right, arr) }
}

func main() {
	var arr [6]int = [6]int{-9, 78, 0, 23, -567, 70}
	fmt.Println(arr)
	QuickSort(0, len(arr) - 1, &arr)
	fmt.Println(arr)
}
