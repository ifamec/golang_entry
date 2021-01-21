package main

import "fmt"

func main() {
	// 1.
	var arr01 [3]int
	arr01[0] = 1 // √
	//arr01[1] = 3.1 	// ERROR DATA TYPE NEED BT THE SAME
	//arr01[3] = 4		// ERROR OUT OF BOUND LENGTH NOT DYNAMIC
	fmt.Println(arr01)

	// 4.
	var arr021 [3]float64
	var arr022 [3]string
	var arr023 [3]bool
	fmt.Println(arr021, arr022, arr023)

	//8
	var arr08 [3]int = [3]int{11,22,33}
	test(arr08)
	fmt.Println(arr08) //[11 22 33]

	// pass by reference to change the value
	var arr09 [3]int = [3]int{11,22,33}
	testR(&arr09)
	fmt.Printf("%p\t", &arr09)
	fmt.Println(arr09)

}

func test(arr [3]int) {
	arr[0] = 88
}
func testR(arr *[3]int) {
	fmt.Printf("testR: %p\t", &arr)
	(*arr)[0] = 88
}