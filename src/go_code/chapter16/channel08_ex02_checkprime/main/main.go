package main

import (
	"fmt"
	"time"
)
var goroutineNumber int = 4
var testNumber int = 20000000

func createIntChannel(intChan chan int) {
	for i := 1; i <= testNumber; i++ {
		intChan <- i
	}
	close(intChan)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i * i <= n; i ++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func checkPrime(intChan chan int, resChan chan int, exitChan chan bool) {
	for {
		n, hasElement := <- intChan
		if !hasElement {
			exitChan <- true
			return
		}
		if isPrime(n) {
			resChan <- n
		}
	}
}

// intChan -> write raw number use one goroutine close at the end
// 4 channel -> get data from intChan check isPrime if true write number to resChan
func main()  {
	var intChan chan int = make(chan int, 1000)
	var resChan chan int = make(chan int, testNumber)
	var exitChan chan bool = make(chan bool, goroutineNumber)

	start := time.Now().UnixNano()

	go createIntChannel(intChan)
	for i := 0; i < goroutineNumber; i++ {
		go checkPrime(intChan, resChan, exitChan)
	}

	// go func() {
		for { // block - check if all goroutine is done
			if len(exitChan) == goroutineNumber {
				end := time.Now().UnixNano()
				fmt.Println("goroutine Time:", end-start)
				break
			}
		}
		// for i := 0; i < 4; i++ {
		// 	<- exitChan
		// }
		// end := time.Now().Unix()
		// fmt.Println("goroutine Time:", end-start)
		close(resChan)
	// }()

	// for v := range resChan {
	// 	fmt.Printf("%v\n", v)
	// }

}