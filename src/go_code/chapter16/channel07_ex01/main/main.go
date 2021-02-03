package main

import (
	"fmt"
)

var N int = 2000

type VSum struct {
	Key int
	Value int
}

func writeNum(ch chan int) {
	for i := 1; i <=N; i++{
		ch <- i
		// fmt.Println(i)
	}
	close(ch)
}

func calcSum(numChan chan int, resChan chan VSum, exitChan chan bool) {
	for {
		n, isSuccess := <- numChan
		if !isSuccess {
			exitChan <- true
			return
		}
		rst, i := 0, 1
		for ; i <= n; i++ {
			rst += i
		}
		resChan <- VSum{n, rst}
	}
}

func main() {
	var numChan chan int = make(chan int, N)
	var resChan chan VSum = make(chan VSum, N)
	var exitChan chan bool = make(chan bool, 8)
	
	go writeNum(numChan)
	for i := 0; i < 8; i++ {
		go calcSum(numChan, resChan, exitChan)
	}

	for {
		if len(exitChan) == 8 {
			break
		}
	}
	close(resChan)

	for v := range resChan {
		fmt.Println(v.Key, v.Value)
	}

}
