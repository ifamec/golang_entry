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
func (p Phone) Start() {
	fmt.Println("Phone Start Working")
}
func (p Phone) Stop() {
	fmt.Println("Phone Stop Working")
}

type Camera struct {// realise all Usb interface method
	Name string
}
func (c Camera) Start() {
	fmt.Println("Camera Start Working")
}
func (c Camera) Stop() {
	fmt.Println("Camera Stop Working")
}


func main()  {
	// struct variable
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = Phone{"iPhone"}
	usbArr[1] = Phone{"Pixel"}
	usbArr[2] = Camera{"Xperia"}
	fmt.Println(usbArr)
}