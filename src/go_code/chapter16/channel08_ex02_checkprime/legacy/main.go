package main

import (
	"fmt"
	"time"
)
// 60 s
var testNumber int = 20000000

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i * i <= n; i ++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now().UnixNano()

	for i :=2; i <=testNumber; i++ {
		if isPrime(i) {
			// do something
		}
	}

	end := time.Now().UnixNano()
	fmt.Println("goroutine Time:", end-start)
}
