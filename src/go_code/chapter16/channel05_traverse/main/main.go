package main

import "fmt"

func main() {
	// traverse channel
	// USE for-range
	var intChan chan int = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	close(intChan)

	// traverse
	// for i := 0; i < len(intChan), i++ {} Nope
	// if not close and traverse error fatal error: all goroutines are asleep - deadlock!
	for val := range intChan {
		fmt.Println(val)
	}
}
