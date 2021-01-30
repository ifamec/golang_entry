package main

import "fmt"

func getSum(n int) (rtnval int) {
	for i := 1; i <= n; i++ {
		rtnval += i
	}
	return
}

func main() {
	// Legacy Test - use and check the result
	res := getSum(10)
	if res != 55 {
		fmt.Println("Test Fails, Expect 55, Received:", res)
	} else {
		fmt.Println("Test Pass")
	}
}
