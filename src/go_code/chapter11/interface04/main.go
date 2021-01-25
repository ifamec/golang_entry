package main

import "fmt"

type A interface {
	test01()
	test02()
}
type B interface {
	test01()
	test03()
}
type C interface {
	A
	B
}

type Usb interface {
	Say()
}

type Stu struct {

}

func (s *Stu) Say() {
	fmt.Println("Stu Say")
}

func main() {
	var c C
	fmt.Println(c)

	var s Stu = Stu{}
	// var u Usb = s // Error this is a not a pointer
	var u Usb = &s
	u.Say()
	fmt.Println(u)

}
