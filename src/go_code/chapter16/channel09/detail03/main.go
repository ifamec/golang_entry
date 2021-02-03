package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello", i)
	}
}
func test() {
	// use recover to capture panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Test Panic !!!")
		}
	}()
	var m map[int]string
	m[0] = "go"
}

func main() {
	go hello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(time.Second)
	}
}
