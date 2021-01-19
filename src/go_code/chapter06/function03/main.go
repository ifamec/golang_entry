package main

import "fmt"

func test(n1 int)  {
	n1 = n1 + 1
	fmt.Println("test: n1 =", n1)
}
func getSum(n1 int, n2 int) (sum int) {
	sum = n1 + n2
	fmt.Println("test: sum =", sum)
	return
}
func getSumDiff(n1 int, n2 int) (sum int, diff int)  {
	sum = n1 + n2
	diff = n1 - n2
	return
}
func main()  {
	n1 := 10

	test(n1)
	fmt.Println("main: n1 =", n1)

	sum := getSum(1, 2)
	fmt.Println("main: sum =", sum)

	res1, res2 := getSumDiff(10, 20)
	fmt.Println("sum: res1 =", res1, "diff: res2 =", res2)

	// if just one value, use "_" placeholder to ignore unneeded rtnval
	res3, _ := getSumDiff(3, 5)
	fmt.Println("sum res3=", res3)
}