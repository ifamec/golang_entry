package main

import "fmt"

func main() {
	// channel can be read/write only

	// default is RW

	// write only
	var chan2 chan <- int = make(chan int, 3)
	chan2 <- 10
	// <- chan2 // Error
	fmt.Println(chan2)

	// read only
	var chan3 <- chan int = make(chan int, 3)
	// chan3 <- 10 // Error
	<- chan3
	fmt.Println(chan3)
}
