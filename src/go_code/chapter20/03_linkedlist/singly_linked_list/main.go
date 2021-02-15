package main

import (
	"fmt"
)

type PNode struct {
	No   int
	Name string
	next *PNode
}

func Append(head *PNode, node *PNode) {
	// get last node
	tNode := head
	for {
		if tNode.next == nil {
			break
		}
		tNode = tNode.next
	}
	tNode.next = node
}

func InsertByNo(head *PNode, node *PNode) {
	// get appropriate node
	tNode := head
	flag := true
	for {
		if tNode.next == nil {
			break
		} else if tNode.next.No > node.No { // node should insert after this
			break
		} else if tNode.next.No == node.No  {
			flag = false
			break
		}
		tNode = tNode.next
	}
	if !flag {
		fmt.Println("Exist", node)
		return
	} else {
		node.next = tNode.next
		tNode.next = node
	}
}

func Delete(head *PNode, no int) {
	tNode := head
	flag := false
	for {
		if tNode.next == nil {
			break
		} else if tNode.next.No == no  {
			flag = true
			break
		}
		tNode = tNode.next
	}

	if flag {
		tNode.next = tNode.next.next
	} else {
		fmt.Println("Id Not Exist")
	}
}

func List(head *PNode) {
	tNode := head
	if tNode.next == nil { // is empty list
		fmt.Println("List is empty")
		return
	}

	for {
		fmt.Printf("%d %v %p\n", tNode.next.No, tNode.next.Name, tNode.next.next)
		tNode = tNode.next
		if tNode.next == nil {
			break
		}
	}

}

func main() {
	// head node
	head := &PNode{}

	h1 := &PNode{
		No:   1,
		Name: "AAA",
	}

	// insert h1 to head
	Append(head, h1)
	Append(head, &PNode{
		No:   2,
		Name: "BBB",
	})
	Append(head, &PNode{
		No:   4,
		Name: "DDD",
	})
	List(head)


	InsertByNo(head, &PNode{
		No:   3,
		Name: "CCC",
	})
	InsertByNo(head, &PNode{
		No:   3,
		Name: "CCC",
	})
	List(head)

	Delete(head, 5)
	Delete(head, 3)
	List(head)
	Delete(head, 4)
	List(head)
}
