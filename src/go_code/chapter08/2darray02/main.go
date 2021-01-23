package main

import "fmt"

func main() {
	var a [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%d\t", a[i][j])
		}
		fmt.Println()
	}

	for _, row := range a {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}

}
