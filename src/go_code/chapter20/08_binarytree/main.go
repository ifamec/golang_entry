package main

import "fmt"

type BTree struct {
	Id int
	Left *BTree
	Right *BTree
}

func PreOrder(node *BTree) {
	// root left right
	if node != nil {
		fmt.Printf("%d\t", node.Id)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
func InOrder(node *BTree) {
	// left root right
	if node != nil {
		InOrder(node.Left)
		fmt.Printf("%d\t", node.Id)
		InOrder(node.Right)
	}
}
func PostOrder(node *BTree) {
	// left right root
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("%d\t", node.Id)
	}
}

func main()  {
	root := &BTree{
		Id: 1,
	}
	ll := &BTree{
		Id: 2,
	}
	root.Left = ll

	rr := &BTree{
		Id: 3,
	}
	root.Right = rr

	rrl := &BTree{
		Id: 4,
	}
	rr.Left = rrl

	rrr := &BTree{
		Id: 5,
	}
	rr.Right = rrr

	fmt.Println("\nPreOrder")
	PreOrder(root)

	fmt.Println("\nInOrder")
	InOrder(root)

	fmt.Println("\nPostOrder")
	PostOrder(root)
}