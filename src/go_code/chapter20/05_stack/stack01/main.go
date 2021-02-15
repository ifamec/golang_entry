package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	maxSize int    // maxSize
	Top     int    // stack top
	Array   [5]int // mock stack
}

func (s *Stack) Push(n int) (err error) {
	if (*s).Top == (*s).maxSize-1 {
		err = errors.New("stack full")
		return
	}
	(*s).Top++
	(*s).Array[(*s).Top] = n
	return
}

func (s *Stack) Pop() (n int, err error) {
	if (*s).Top == -1 {
		err = errors.New("stack empty")
		return
	}
	n = (*s).Array[(*s).Top]
	(*s).Top--
	return
}

func (s *Stack) List() (err error) {
	if (*s).Top == -1 {
		err = errors.New("stack empty")
		return
	}
	fmt.Println("List Stack")
	for i := (*s).Top; i >= 0; i-- {
		fmt.Println(i, (*s).Array[i])
	}
	return
}

func main() {
	var stack = &Stack{
		maxSize: 5,
		Top:     -1, // empty
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.List()

	n, err := stack.Pop()
	if err!= nil {
		fmt.Println("Error -", err)
	}
	fmt.Println(n)
	stack.List()
}
