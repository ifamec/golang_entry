package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // head is the first element
	tail    int // tail is not the last element
}

func (c *CircleQueue) Push(n int) (err error) {
	// check if array is full
	if c.IsFull() {
		err = errors.New("add queue fail, queue is full")
		return
	}
	c.array[c.tail] = n
	c.tail = (c.tail + 1) % c.maxSize

	return
}

func (c *CircleQueue) Pop() (val int, err error) {
	// check if array is empty
	if c.IsEmpty() {
		err = errors.New("queue is empty")
		return
	}
	val = c.array[c.head]
	c.head =  (c.head + 1) % c.maxSize

	return
}

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

func (c *CircleQueue) ListQueue() {
	size := c.Size()
	if size == 0 {
		fmt.Println("Queue is empty")
		return
	}
	// find head then traverse to tail
	fmt.Println("Current Queue:")
	var tHead int = c.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d] = %d\n", tHead, c.array[tHead])
		tHead = (tHead + 1) % c.maxSize
	}
	fmt.Println()
}

func main() {
	cQueue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
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

			err := cQueue.Push(val)
			if err != nil {
				fmt.Println("Error -", err)
			} else {
				fmt.Println("Add Queue Success")
			}
		case "get":
			val, err := cQueue.Pop()
			if err != nil {
				fmt.Println("Error -", err)
			} else {
				fmt.Println("Get Queue Success:", val)
			}
		case "list":
			cQueue.ListQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Retry")
		}
	}

}
