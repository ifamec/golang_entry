package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}
func (p *Emp) ShowDetail() {
	fmt.Printf("Index: %d, Id: %d, Name: %s\n", p.Id % 8 ,p.Id, p.Name)
}

type EmpLink struct {
	Head *Emp
}
func (e *EmpLink) Insert(emp *Emp) {
	// sm -> lg
	cur := e.Head      // temp1
	var pre *Emp = nil // temp2 always ahead of cur
	if cur == nil {    // empty
		e.Head = emp
		return
	}
	for { // find emp position
		if cur != nil {
			if cur.Id > emp.Id { // insert before cur, after pre
				// insert
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	// insert after pre
	pre.Next = emp
	emp.Next = cur
}
func (e *EmpLink) List(num int) {
	fmt.Printf("%d->", num)
	if e.Head == nil {
		fmt.Printf("nil\n")
		return
	}
	cur := e.Head // temp
	for {
		if cur != nil {
			fmt.Printf("(id:%v name:%v)->", cur.Id, cur.Name)
			cur = cur.Next
		} else {
			fmt.Printf("nil\n")
			break
		}
	}
}
func (e *EmpLink) Find(num int) *Emp {
	cur := e.Head // temp
	for {
		if cur != nil && num == cur.Id {
			return cur
		} else if cur == nil {
			return nil
		}
		cur = cur.Next
	}
}

type HashTable struct {
	LinkArr [8]EmpLink
}

func (h *HashTable) InsertEmp(emp *Emp) {
	// which linked list
	linkNum := h.HashFn(emp.Id)
	h.LinkArr[linkNum].Insert(emp)
}
func (h *HashTable) HashFn(id int) int {
	return id % 8
}
func (h *HashTable) ListAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].List(i)
	}
}
func (h *HashTable) FindById(id int) *Emp {
	return h.LinkArr[h.HashFn(id)].Find(id)
}

func main() {
	var key string
	var id int
	var name string
	var ht HashTable
	for {
		fmt.Println("===== Select =====")
		fmt.Println("input")
		fmt.Println("list")
		fmt.Println("find")
		fmt.Println("exit")
		fmt.Print("===== Select: ")
		fmt.Scanln(&key)

		switch key {
		case "input":
			fmt.Printf("id: ")
			fmt.Scanln(&id)
			fmt.Printf("name: ")
			fmt.Scanln(&name)
			ht.InsertEmp(&Emp{
				Id: id,
				Name: name,
			})
		case "list":
			ht.ListAll()
		case "find":
			fmt.Printf("id: ")
			fmt.Scanln(&id)
			emp := ht.FindById(id)
			if emp == nil {
				fmt.Println(id, "Not Found")
			} else {
				emp.ShowDetail()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Retry")
		}
	}
}
