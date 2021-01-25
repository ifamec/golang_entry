package main

import "fmt"

type Usb interface {
	// declare two unrealised method
	Start()
	Stop()
}

type Phone struct {// realise all Usb interface method

}
func (p *Phone) Start() {
	fmt.Println("Phone Start Working")
}
func (p *Phone) Stop() {
	fmt.Println("Phone Stop Working")
}

type Camera struct {// realise all Usb interface method

}
func (c *Camera) Start() {
	fmt.Println("Camera Start Working")
}
func (c *Camera) Stop() {
	fmt.Println("Camera Stop Working")
}
type PC struct {

}
func (pc *PC) Working(usb Usb) { // receive Usb interface data type
	// use interface to use the method
	usb.Start()
	usb.Stop()
}

func main()  {
	// struct variable
	var pc = &PC{}
	var phone = &Phone{}
	var camera = &Camera{}

	pc.Working(phone)
	pc.Working(camera)
}