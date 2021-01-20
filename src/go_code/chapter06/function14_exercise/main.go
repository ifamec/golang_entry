package main

import "fmt"

func printPyramid(layer int)  {
	// full pyramid
	for i := 1; i <= layer; i ++ {
		for j := 1; j <= layer - i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func printTimesTable(number int)  {
	for i := 1; i <=number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}

func conversion(arr [3][3] int) (rtnval [3][3] int) {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			rtnval[i][j] = arr[j][i]
		}
	}
	return
}

func main()  {
	//var layer, number int
	//fmt.Println("Input Layer:")
	//fmt.Scanf("%d\n", &layer)
	//printPyramid(layer)
	//
	//fmt.Println("Input Number:")
	//fmt.Scanf("%d\n", &number)
	//printTimesTable(number)

	arr := [3][3] int {
		{1,2,3}, {4,5,6}, {7,8,9},
	}
	fmt.Println(arr)
	fmt.Println(conversion(arr))
}