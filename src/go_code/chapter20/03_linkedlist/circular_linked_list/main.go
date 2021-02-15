package main

import (
	"fmt"
)

type CatNode struct {
	No   int
	Name string
	next *CatNode
}

func Append(head *CatNode, node *CatNode) {
	// if first cat
	if head.next == nil {
		head.No = node.No
		head.Name = node.Name
		head.next = head
		return
	}
	// get last node
	tNode := head
	for {
		if tNode.next == head {
			break
		}
		tNode = tNode.next
	}
	tNode.next = node
	node.next = head
}

//
// func InsertByNo(head *CatNode, node *CatNode) {
// 	// get appropriate node
// 	tNode := head
// 	flag := true
// 	for {
// 		if tNode.next == nil {
// 			break
// 		} else if tNode.next.No > node.No { // node should insert after this
// 			break
// 		} else if tNode.next.No == node.No  {
// 			flag = false
// 			break
// 		}
// 		tNode = tNode.next
// 	}
// 	if !flag {
// 		fmt.Println("Exist", node)
// 		return
// 	} else {
// 		node.next = tNode.next
// 		tNode.next = node
// 	}
// }

func Delete(head *CatNode, no int) *CatNode {
	tNode := head
	hNode := head
	flag := true
	if tNode.next == nil {
		fmt.Println("Empty")
		return head
	}
	if tNode.next == head {
		fmt.Println("Only One Node")
		tNode.next = nil
		return head
	}
	for {
		if hNode.next == head {
			break
		}
		hNode = hNode.next
	}
	for {
		if tNode.next == head { // did not find BEFORE last one
			break
		}
		if tNode.No == no {
			if tNode == head {
				head = head.next
			}
			hNode.next = tNode.next
			fmt.Println(no, "Get Deleted")
			flag = false
			break
		}
		tNode = tNode.next // for compare
		hNode = hNode.next // for delete
	}
	// compare last
	if flag {
		if tNode.No == no {
			hNode.next = tNode.next
			fmt.Println(no, "Get Deleted")
		} else {
			fmt.Println(no, "Not Exist")
		}
	}

	return head
}

func List(head *CatNode) {
	tNode := head
	if tNode.next == nil { // is empty list
		fmt.Println("List is empty")
		return
	}

	for {
		fmt.Printf("%d %v %p\n", tNode.No, tNode.Name, tNode.next)
		if tNode.next == head {
			break
		}
		tNode = tNode.next
	}

}

func main() {
	// head node
	head := &CatNode{}

	c1 := &CatNode{
		No:   1,
		Name: "AAA",
	}

	// insert h1 to head
	Append(head, c1)
	Append(head, &CatNode{
		No:   2,
		Name: "BBB",
	})
	Append(head, &CatNode{
		No:   4,
		Name: "DDD",
	})
	List(head)

	//
	// InsertByNo(head, &CatNode{
	// 	No:   3,
	// 	Name: "CCC",
	// })
	// InsertByNo(head, &CatNode{
	// 	No:   3,
	// 	Name: "CCC",
	// })
	// List(head)

	fmt.Println()
	head = Delete(head, 5)
	List(head)
	fmt.Println()
	head = Delete(head, 4)
	List(head)
	fmt.Println()
	head = Delete(head, 1)
	List(head)
	fmt.Println()
	head = Delete(head, 2)
	List(head)
	fmt.Println()
	head = Delete(head, 1)
	List(head)
}
