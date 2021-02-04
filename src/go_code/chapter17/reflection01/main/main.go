package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func (s *Student) Say() {
	fmt.Println("Hello")
}

func rInt(b interface{}) {
	// 1. get reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	// 2. get reflect.Value
	rValue := reflect.ValueOf(b)
	// var n1 int = rValue // Cannot use 'rValue' (type Value) as type int
	fmt.Printf("Value: %v, Type: %T\n", rValue, rValue)

	// get real value
	n2 := 1 + rValue.Int()
	fmt.Println(n2)

	// 3. to Interface
	iValue := rValue.Interface()
	fmt.Println("interface type assertion:", iValue.(int))
}

func rStruct(b interface{}) {
	// 1. get reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	// 2. get reflect.Value
	rValue := reflect.ValueOf(b)
	// fmt.Printf("Value: %v, Type: %T\n", rValue, rValue)

	// 2.5 get reflect.Value.Kind
	fmt.Println("rType.Kind():", rType.Kind())
	fmt.Println("rValue.Kind():", rValue.Kind())

	// 3. to Interface
	iValue := rValue.Interface()
	fmt.Printf("to interface: %v, %T\n", iValue, iValue)
	// fmt.Println(iValue.Name) // Error Decide when executing the code
	// fmt.Println(iValue.(Student).Name)
	// type assertion with check
	stu, ok := iValue.(Student)
	if ok {
		fmt.Println(stu.Name)
	}

}

func main() {
	var num int = 100
	// get Type Kind use reflection
	rInt(num)

	fmt.Println("\n===\nStruct:")
	var stu Student = Student{"Tom", 20}
	rStruct(stu)
}
