package main

import "fmt"

func main()  {
	var intChan chan int = make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// intChan <- 300 // panic: send on closed channel
	fmt.Println("Could Read", <- intChan)
	fmt.Println("Could Read", <- intChan)
}