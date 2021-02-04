package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func rInt(b interface{}) {

	// 2. get reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("%v\n", rValue.Kind())
	// rValue.SetInt(20) // panic: reflect: reflect.Value.SetInt using unaddressable value
	rValue.Elem().SetInt(20) // Elem returns the value that the interface v contains or that the pointer v points to.
	fmt.Printf("Value: %v, Type: %T\n", rValue, rValue)

}

func rStruct(b interface{}) {

	// 2. get reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue.Kind():", rValue.Kind())

	actualStruct := rValue.Elem()
	fmt.Printf("%T, %v, %v\n", actualStruct, actualStruct.Type(), actualStruct.Kind())
	field := actualStruct.FieldByName("Name")
	fmt.Printf("%T, %v, %v\n", field, field.Type(), field.Kind())
	field.SetString("BBB")
}

func main() {
	// use reflection, modify the value of num
	var num int = 100
	rInt(&num)
	fmt.Println(num)

	// modify student value
	fmt.Println("\n===\nStruct:")
	stu := &Student{"Tom", 20}
	rStruct(stu)
	fmt.Println(*stu)
}
