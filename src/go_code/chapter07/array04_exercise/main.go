package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	ex01()
	ex02()
	ex03()
	ex04()
}

func ex01()  {
	var alphabet [26]byte
	for i := 0; i < 26; i++ {
		alphabet[i] = 'A' + byte(i)
	}
	for _, val := range alphabet {
		fmt.Printf("%c  ", val)
	}
	fmt.Println("")
}

func ex02()  {
	var arr = [...]int{1,4,-6,3,8,7}
	var max int = arr[0]
	var idx int = 0
	for index, value := range arr {
		if max < value {
			max = value
			idx = index
		}
	}
	fmt.Println(idx, max)
}

func ex03() {
	var arr = [...]int{1,4,-6,11,2,3,8,7,5}
	var sum int
	for _, value := range arr {
		sum += value
	}
	fmt.Println("sum:", sum, "avg:", float64(sum)/float64(len(arr)))
}

func ex04()  {
	var arr [5]int

	rand.Seed(time.Now().UnixNano())
	for index, _ := range arr{
		arr[index] = rand.Intn(100)
	}
	fmt.Println(arr)
	arrLen := len(arr)
	for i := 0; i < arrLen/2; i++ {
		arr[i], arr[arrLen-1-i] = arr[arrLen-1-i], arr[i]
	}
	fmt.Println(arr)
}