package main

import "fmt"

func main()  {
	test1(4)
	test2(4)
	fmt.Println("Fibonacci: ", Fibonacci(5))
	fmt.Println("solveIssue: ", solveIssue(5))
	fmt.Println("peach: ", peach(0))
}

func test1(n int)  {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println("test1: n =", n)
}
/* STACK
test1 n=2 -> Print n = 2 (1) pop (1)
test1 n=3 -> n=2 | -> Print n = 2 (2) pop (2)
test1 n=4 -> n=3 | -> Print n = 3 (3) pop (3)
main test1(4)
*/
func test2(n int)  {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("test2: n =", n)
	}
}
/* STACK

test2 n=2 -> Print n = 2 pop (1)
test2 n=3 -> n=2 | -> nil never goto else pop (2)
test2 n=4 -> n=3 | -> nil never goto else pop (3)
main test2(4)
*/
func Fibonacci(n int) int {
	// when n == 1 || n == 2
	// when n > 2 return sum of prev two num
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
func solveIssue(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2 * solveIssue(n - 1) + 1
	}
}
func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("ERROR")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return 2 * (peach(n + 1) + 1)
	}
}