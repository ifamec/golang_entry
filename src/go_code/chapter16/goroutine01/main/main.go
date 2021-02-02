package main

import (
	"fmt"
	"time"
)

func printStr() {
	for i := 1; i <= 10; i++ {
		fmt.Println("goroutine,", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go printStr() // start a goroutine
	for i := 1; i <= 10; i++ {
		fmt.Println("mainThread,", i)
		time.Sleep(time.Second)
	}
}
