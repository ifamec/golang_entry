package main

import "fmt"

type A interface {
	Say()
}
type B interface {
	Hello()
}
type C interface {
	A
	B
	Oh()
}

type realiseInterfaceC struct {

}

func (r *realiseInterfaceC) Say() {
	fmt.Println("realiseInterfaceC - Say")
}
func (r *realiseInterfaceC) Hello() {
	fmt.Println("realiseInterfaceC - Hello")
}
func (r *realiseInterfaceC) Oh() {
	fmt.Println("realiseInterfaceC - Oh")
}

type Empty interface {}

func main()  {
	var c realiseInterfaceC
	var ic C = &c
	c.Say()
	c.Hello()
	c.Oh()

	ic.Oh()

	var e Empty = &c
	var e2 interface{} = 10
	fmt.Println(e, e2)

}