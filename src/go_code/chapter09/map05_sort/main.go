package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[int]int = map[int]int{
		10: 2,
		2: 67,
		5: 3,
		7: 93,
		1: 36,
		4: 73,
		9: 40,
	}
	for key, value := range m {
		fmt.Printf("%v: %v\t", key, value)
	}
	fmt.Println()
	// sort by key
	// key to slice
	// sort slice
	// traverse slice and print value
	var keys []int
	for key, _ := range m {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Printf("m[%v]=%v\t", key, m[key])
	}
}
