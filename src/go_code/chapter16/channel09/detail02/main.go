package main

import "fmt"

func main() {
	var intChan chan int = make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	var strChan chan string = make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "Hello" + fmt.Sprintf("%d", i)
	}

	// Legacy way to traverse if not close, will cause deadlock
	// ISSUE: kind of hard to estimate when to close the channel

	// label:
	for {
		select {
		case v := <-intChan: // if channel not closed, no deadlock, code will continue
			fmt.Println("intChan Data:", v)
		case v := <-strChan:
			fmt.Println("strChan Data:", v)
		default:
			fmt.Println("Done")
			// break label
			return
		}
	}
}
