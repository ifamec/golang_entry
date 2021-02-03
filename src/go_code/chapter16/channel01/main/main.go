package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("%v, %p\n", intChan, &intChan) // Address | Reference Type

	var num = 11
	// write data
	intChan <- 10
	intChan <- num
	intChan <- 12
	// intChan <- 13 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println("len", len(intChan), "cap", cap(intChan)) // cap will not increase automatically

	// read
	var num2 int
	num2 = <- intChan
	fmt.Println(num2)
	fmt.Println("len", len(intChan), "cap", cap(intChan)) // cannot pop to much if not using goroutine // Error

	<- intChan
	<- intChan
	fmt.Println("len", len(intChan), "cap", cap(intChan)) // cannot pop to much if not using goroutine // Error

	<- intChan // fatal error: all goroutines are asleep - deadlock!
}
