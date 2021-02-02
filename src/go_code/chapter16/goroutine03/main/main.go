package main

import (
	"fmt"
	"sync"
	"time"
)

// declare a mutex lock
var lock = sync.Mutex{}

// Map
var Result = make(map[int]uint64, 10)

func Factorial(n int) {
	var rtnval uint64 = 1
	for i := 1; i <= n; i++ {
		rtnval *= uint64(i)
	}
	// put into map
	lock.Lock()
	Result[n] = rtnval // TODO fatal error: concurrent map writes USE LOCK
	lock.Unlock()
}

func main() {
	// factorial function
	// multiple goroutine
	// map should in global scope
	// start go routine
	for i := 1; i<=30; i++ {
		go Factorial(i)
	}
	// sleep 10s
	time.Sleep(10 * time.Second) // TODO WAIT HOW MANY TIME
	// result
	lock.Lock()
	for i, v := range Result {
		fmt.Println(i, v)
	}
	lock.Unlock()
}
