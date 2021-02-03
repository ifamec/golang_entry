package main

import (
	"fmt"
	"time"
)

// writeData goroutine, write 50 nums into intChannel
// readData goroutine, read data from intChannel
// use same channel and main thread need to wait

// thought: when read all data write bool to a exitChannel,
// in main thread use for loop to read exit channel if true then exit

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		// time.Sleep(time.Second)
		fmt.Println("WRITE:", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		val, isSuccess := <-intChan
		if !isSuccess {
			break
		}
		time.Sleep(time.Millisecond * 500)
		fmt.Println("READ:", val)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	var intChan chan int = make(chan int, 50)
	var exitChan chan bool = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, isSuccess := <- exitChan
		if !isSuccess {
			break
		}
	}

}
