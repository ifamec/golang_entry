package main

import "fmt"

func main() {
	var arr [4][6]int
	fmt.Println(arr)
	arr[0] = [6]int{0, 0, 0, 0, 0, 0}
	arr[1] = [6]int{1, 1, 1, 1, 1, 1}
	arr[2][2] = 2
	arr[2][4] = 3
	for _, row := range arr {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	arr = [4][6]int{{1, 1, 1, 1, 1, 1}, {2, 2, 2, 2, 2, 2}, {3, 3, 3, 3, 3, 3}, {4, 4, 4, 4, 4, 4}}
	fmt.Println(arr)

	var array [2][3]int
	array[1][1] = 2
	fmt.Printf("%p\t%p\n%p\t%p\n", &array[0], &array[0][0], &array[1], &array[1][0])

	var a [2][3]int = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(a)
}
