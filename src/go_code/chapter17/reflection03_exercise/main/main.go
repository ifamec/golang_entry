package main

import (
	"fmt"
	"reflect"
)

func rFloat(b interface{}) {
	rValue := reflect.ValueOf(b)

	fmt.Println("Kind", rValue.Kind())
	fmt.Println("Type", rValue.Type())

	iValue := rValue.Interface()

	var value float64 = iValue.(float64)
	fmt.Printf("%v, %T\n", value, value)
}

func main() {
	var v float64 = 1.2
	rFloat(v)

	var str string = "Tom"
	fs := reflect.ValueOf(&str)
	// fs := reflect.ValueOf(str) // need to be an address
	// fs.SetString("Jerry") // panic: reflect: reflect.Value.SetString using unaddressable value
	fs.Elem().SetString("Jerry")
	fmt.Println(str)
}
