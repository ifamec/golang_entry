package main

import "fmt"

type A interface {
	Say()
}

type Stu struct {
	Name string
}

func (s *Stu) Say() {
	fmt.Println("STU SAY:", (*s).Name)
}

type integer int

func (i *integer) Say() {
	fmt.Println("integer SAY:", *i)
}

type B interface {
	Hello()
}
type Ms struct {

}

func (m *Ms) Hello() {
	fmt.Println("Ms Hello")
}
func (m *Ms) Say() {
	fmt.Println("Ms Say")
}


func main()  {
	// var a A
	// a.Say() // panic: runtime error: invalid memory address or nil pointer dereference

	var s = &Stu{"Tom"} // Realise the A interface
	var a A = s
	a.Say()

	var i integer = 10
	var b A = &i
	b.Say()

	var m0 Ms
	var mA A = &m0
	var mB B = &m0
	m0.Say(); m0.Hello()
	mA.Say()
	mB.Hello()

}