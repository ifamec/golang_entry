package main

import (
	"fmt"
)

func main()  {
	var arr [3] string = [3] string {"AAA", "BBB", "CCC"}
	for index, value := range arr{
		fmt.Println(index, value, arr[index])
	}

	for _, value := range arr{ // value only
		fmt.Printf("%v\t", value)
	}
}