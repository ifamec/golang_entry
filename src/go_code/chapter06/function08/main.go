package main

import "fmt"

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func main()  {
	var i, j int = 2, 3
	fmt.Println(i, j)
	swap(&i, &j)
	fmt.Println(i, j)
}