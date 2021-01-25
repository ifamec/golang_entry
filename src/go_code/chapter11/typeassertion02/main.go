package main

import "fmt"

type Usb interface {
	// declare two unrealised method
	Start()
	Stop()
}

type Phone struct {// realise all Usb interface method
	Name string
}
func (p *Phone) Start() {
	fmt.Println("Phone Start Working")
}
func (p *Phone) Stop() {
	fmt.Println("Phone Stop Working")
}
func (p *Phone) Call() {
	fmt.Println("Phone Calling")
}

type Camera struct {// realise all Usb interface method
	Name string
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
	// if usb pointing to Phone, Call()
	// type assertion
	if phone, flag := usb.(*Phone); flag {
		phone.Call()
	}
	usb.Stop()
}

func main()  {
	// struct variable
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = &Phone{"iPhone"}
	usbArr[1] = &Phone{"Pixel"}
	usbArr[2] = &Camera{"Alpha"}
	fmt.Println(usbArr[0], usbArr[1], usbArr[2])

	var pc = &PC{}
	for _, v:= range usbArr {
		pc.Working(v)
		fmt.Println()
	}
}