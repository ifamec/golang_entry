package main

import "fmt"

func Fibonacci(n int)  {
	var fibo []int
	if n == 1 {
		fibo = append(fibo, 1)
	} else if n == 2 {
		fibo = append(fibo, 1, 1)
	} else {
		fibo = append(fibo, 1, 1)
		for i := 2; i < n; i++ {
			fibo = append(fibo, fibo[i-2] + fibo[i-1])
		}
	}
	fmt.Println(fibo)
}

func Fibo(n int) (fibo []uint64) {
	fibo = make([]uint64, n)
	fibo[0] = 1
	fibo[1] = 1
	for i := 2; i < n; i++ {
		fibo[i] = fibo[i-2] + fibo[i-1]
	}
	return
}

func main()  {
	// Fibonacci function
	// get int
	// put Fibonacci into slice
	Fibonacci(1)
	Fibonacci(2)
	Fibonacci(3)
	Fibonacci(4)
	Fibonacci(5)
	fmt.Println()

	fmt.Println(Fibo(100))
}