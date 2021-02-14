package main

import (
	"errors"
	"fmt"
	"os"
)

type ArrayQueue struct {
	maxSize int
	array [5]int
	front int // front + 1 point to head of queue
	rear int
}

func (a *ArrayQueue) addQueue(n int) (err error) {
	// check if array is full
	if a.rear == a.maxSize - 1 {
		err = errors.New("add queue fail, queue is full")
		return
	}
	a.rear ++
	a.array[a.rear] = n

	return
}

func (a *ArrayQueue) getQueue() (val int, err error) {
	// check if array is full
	if a.rear == a.front {
		err = errors.New("queue is empty")
		return
	}
	a.front ++
	val = a.array[a.front]

	return
}

func (a *ArrayQueue) listQueue() {
	// find front then traverse to rear
	fmt.Println("Current Queue:")
	for i := a.front + 1; i <= a.rear; i++ {
		fmt.Printf("%d\t", a.array[i])
	}
	fmt.Println()
}

func main() {
	aQueue := ArrayQueue{
		maxSize: 5,
		front: -1,
		rear: -1,
	}

	var key string
	var val int
	for {
		fmt.Println("==========")
		fmt.Println("1. add to queue")
		fmt.Println("2. get from queue")
		fmt.Println("3. list current queue")
		fmt.Println("4. exit")
		fmt.Print("Select: ")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Print("Number: ")
			fmt.Scanln(&val)

			err := aQueue.addQueue(val)
			if err != nil {
				fmt.Println("Error -", err)
			} else {
				fmt.Println("Add Queue Success")
			}
		case "get":
			val, err := aQueue.getQueue()
			if err != nil {
				fmt.Println("Error -", err)
			} else {
				fmt.Println("Get Queue Success:", val)
			}
		case "list":
			aQueue.listQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Retry")
		}
	}

}
