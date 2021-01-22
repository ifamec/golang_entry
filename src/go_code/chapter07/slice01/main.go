package main

import "fmt"

func main()  {
	var arr [5]int = [...]int{11,33,55,77,99}
	var sli []int = arr[1:3] // index [1,3)

	fmt.Println(sli, len(sli), cap(sli)) // capacity of slice is dynamic
	// generally capacity is 2 * length
	fmt.Printf("%v %v\n", &arr[1], &sli[0])

	sli[1] = 56
	fmt.Println(sli, arr)
}