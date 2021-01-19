package main

import "fmt"

type myFunType func (int, int) int

func test1() {
	//var n1 int = 10 // local variable
}
func test2(n1 int) {
	n1++
	fmt.Println("test2: n1=", n1)
}
func test3(n1 *int) {
	*n1++
	fmt.Println("test3: n1=", *n1)
}
func getSum(n1 int, n2 int) int {
	return n1 + n2
}
func test4(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func test5(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func getSumDiff(n1 int, n2 int) (sum int, diff int)  {
	sum = n1 + n2
	diff = n1 - n2
	return
}
func main() {
	//fmt.Println(n1) // cannot fetch
	n1 := 10
	test2(n1)
	fmt.Println("main: n1=", n1)
	test3(&n1)
	fmt.Println("main: n1=", n1, &n1)

	sum := getSum
	fmt.Printf("type of sum: %T, type of getSum: %T\n", sum, getSum)
	fmt.Printf("sum res: %T, getSum res: %T\n", sum(1,1), getSum(1,1))

	fmt.Println("func as param sum: ", test4(getSum, 1, 2))


	type myInt int // alias int as myInt BUT go think myInt and int are different data type
	var num myInt = 1
	fmt.Printf("%T, %d\n", num, num)
	var nun int = 1
	//num = nun
	nun = int(num)
	fmt.Printf("%T, %d\n", nun, nun)

	fmt.Println("myFunType sum: ", test5(getSum, 1, 2))

	resSum, resDiff := getSumDiff(99, 88)
	fmt.Println("Get Sum and Diff", resSum, resDiff)
}
