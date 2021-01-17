package main

import "fmt"

func main()  {
	var i int = 10
	var j int = 20
	fmt.Println("address:", &i)

	var ptr *int = &i
	fmt.Println("get val from ptr:", *ptr)
	fmt.Printf("\nEXERCISE\n")
	exercise()
	// Ternary

	//var t = i > j ? i : j
	var t int
	if i > j {
		t = i
	} else {
		t = j
	}
	fmt.Println(t)
}

func exercise()  {
	// print max in 2 / 3 vars
	var m int = 20
	var n int = 10
	var o int = 30

	var max_mn int
	if m > n {
		max_mn = m
	} else {
		max_mn = n
	}
	fmt.Println(max_mn)

	var max_mno int
	if o > max_mn {
		max_mno = o
	} else {
		max_mno = max_mn
	}
	fmt.Println(max_mno)
	fmt.Println("")
}