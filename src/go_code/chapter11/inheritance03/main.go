package main

import "fmt"

type A struct {
	Name string
	age int
}

func (a *A) SayOK() {
	fmt.Println("A SayOK", (*a).Name)
}
func (a *A) hello() {
	fmt.Println("A hello", (*a).Name)
}

type B struct {
	A
	Name string
}
func (b *B) SayOK() {
	fmt.Println("B SayOK", (*b).Name)
}

func main()  {
	// var b = &B{}
	// (*b).A.Name = "Tom"
	// (*b).A.age = 19
	// (*b).A.SayOK()
	// (*b).A.hello()
	//
	// (*b).Name = "Jerry"
	// (*b).age = 20
	// (*b).SayOK()
	// (*b).hello()

	var b = &B{}
	(*b).Name = "Jack"
	(*b).A.Name = "Tony"
	(*b).age = 30
	(*b).SayOK()
	(*b).A.SayOK()
	(*b).hello()


}