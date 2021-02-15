package main

import "fmt"

type Object struct {
	No   int
	next *Object
}

// 1. make singly circular linked list
func MakeList(num int) (head *Object) {
	head = &Object{}
	curObj := &Object{}
	if num < 1 {
		return
	}

	for i := 1; i <= num; i++ {
		o := &Object{
			No: i,
		}
		if i == 1 {
			head = o
			curObj = o
			curObj.next = head
		} else {
			curObj.next = o
			curObj = o
			curObj.next = head
		}
	}
	return
}

func List(head *Object) {
	if head.next == nil {
		fmt.Println("Empty List")
		return
	}

	curObj := head
	for {
		fmt.Printf("%v\t", curObj.No)
		if curObj.next == head {
			break
		}
		curObj = curObj.next
	}
}

func Josephus(head *Object, start int, count int) { // leave last obj
	if head.next == nil {
		fmt.Println("Empty List")
		return
	}

	tail := head
	for {
		if tail.next == head {
			break
		}
		tail = tail.next
	}

	// use head
	// move to start
	for i := 0; i < start - 1; i++ {
		head = head.next
		tail = tail.next
	}

	// move count times then delete head
	for {
		// move count - 1 times
		for i := 0; i < count - 1; i++ {
			head = head.next
			tail = tail.next
		}
		fmt.Printf("%d\t", head.No)
		// delete head
		head = head.next
		tail.next = head

		if head == tail {
			break
		}
	}
	fmt.Printf("%d\t", head.No)
	fmt.Println()

}

func main() {
	head := MakeList(500)
	List(head)
	fmt.Println()
	fmt.Println()
	Josephus(head, 20, 31)
}
